package store

import (
	"database/sql"
	"encoding/json"

	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/redis/go-redis/v9"
)

type Postgres struct {
	db     *sql.DB
	secret string
	cache  *redis.Client
}

func NewPostgres(db *sql.DB, secret string) *Postgres {
	return &Postgres{db: db, secret: secret}
}

func (p *Postgres) SetRedis(cache *redis.Client) {
	p.cache = cache
}

func (p *Postgres) Health() error {
	return p.db.Ping()
}

func (p *Postgres) EnsureSchema() error {
	for _, stmt := range schema {
		if _, err := p.db.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func one(db queryer, q string, args ...any) (map[string]any, error) {
	var raw []byte
	if err := db.QueryRow(q, args...).Scan(&raw); err != nil {
		return nil, err
	}
	var out map[string]any
	return out, json.Unmarshal(raw, &out)
}

func many(db queryer, q string, args ...any) ([]map[string]any, error) {
	rows, err := db.Query(q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []map[string]any{}
	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			return nil, err
		}
		var item map[string]any
		if err := json.Unmarshal(raw, &item); err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, rows.Err()
}

func noRow(err error, missing api.Error) api.Error {
	if err == sql.ErrNoRows {
		return missing
	}
	if err != nil {
		return api.ErrInternal
	}
	return api.NoError
}

type queryer interface {
	Query(string, ...any) (*sql.Rows, error)
	QueryRow(string, ...any) *sql.Row
}
