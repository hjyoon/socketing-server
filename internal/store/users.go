package store

import (
	"database/sql"

	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/auth"
)

func (p *Postgres) CreateUser(in Register) (map[string]any, api.Error) {
	if in.Role == "" {
		in.Role = "user"
	}
	var exists int
	_ = p.db.QueryRow(`SELECT count(*) FROM "user" WHERE email=$1`, in.Email).Scan(&exists)
	if exists > 0 {
		return nil, api.ErrUserExists
	}
	salt := auth.NewSalt()
	nick := "user-" + salt[:8]
	pass := auth.HashPassword(in.Password, salt)
	q := `INSERT INTO "user"(email,nickname,password,salt,role)
VALUES($1,$2,$3,$4,$5) RETURNING ` + userJSON
	data, err := one(p.db, q, in.Email, nick, pass, salt, in.Role)
	return data, noRow(err, api.ErrInternal)
}

func (p *Postgres) Login(in Login, secret string) (map[string]any, api.Error) {
	var id, pass, salt string
	err := p.db.QueryRow(`SELECT id,password,salt FROM "user" WHERE email=$1`, in.Email).
		Scan(&id, &pass, &salt)
	if err == sql.ErrNoRows || auth.HashPassword(in.Password, salt) != pass {
		return nil, api.ErrCredentials
	}
	if err != nil {
		return nil, api.ErrInternal
	}
	token := auth.Sign(id, secret)
	return map[string]any{"tokenType": "Bearer", "expiresIn": 86400, "accessToken": token}, api.NoError
}

func (p *Postgres) GetUser(id string) (map[string]any, api.Error) {
	data, err := one(p.db, `SELECT `+userJSON+` FROM "user" WHERE id=$1`, id)
	return data, noRow(err, api.ErrUserNotFound)
}

func (p *Postgres) GetUserByEmail(email string) (map[string]any, api.Error) {
	data, err := one(p.db, `SELECT `+userJSON+` FROM "user" WHERE email=$1`, email)
	return data, noRow(err, api.ErrUserNotFound)
}

func (p *Postgres) GetPoints(id string) (map[string]any, api.Error) {
	data, err := one(p.db, `SELECT json_build_object('id',id,'point',point) FROM "user" WHERE id=$1`, id)
	return data, noRow(err, api.ErrUserNotFound)
}

func (p *Postgres) UpdateNickname(id, nickname string) (map[string]any, api.Error) {
	q := `UPDATE "user" SET nickname=$1,"updatedAt"=now() WHERE id=$2
RETURNING json_build_object('id',id,'nickname',nickname)`
	data, err := one(p.db, q, nickname, id)
	if err == nil {
		return data, api.NoError
	}
	return nil, api.ErrNicknameExists
}

func (p *Postgres) UpdatePassword(id, password string) api.Error {
	var salt string
	if err := p.db.QueryRow(`SELECT salt FROM "user" WHERE id=$1`, id).Scan(&salt); err != nil {
		return noRow(err, api.ErrUserNotFound)
	}
	_, err := p.db.Exec(`UPDATE "user" SET password=$1,"updatedAt"=now() WHERE id=$2`,
		auth.HashPassword(password, salt), id)
	return noRow(err, api.ErrUserNotFound)
}

func (p *Postgres) DeleteUser(id string) api.Error {
	res, err := p.db.Exec(`UPDATE "user" SET "deletedAt"=now() WHERE id=$1`, id)
	if err != nil {
		return api.ErrInternal
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return api.ErrUserNotFound
	}
	return api.NoError
}

const userJSON = `json_build_object('id',id,'nickname',nickname,'email',email,
'profileImage',"profileImage",'role',role,'createdAt',"createdAt",'updatedAt',"updatedAt")`
