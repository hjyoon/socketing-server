package store

import (
	"database/sql"

	"github.com/hjyoon/socketing-server/internal/api"
)

func (p *Postgres) ListEvents(userID string) ([]map[string]any, api.Error) {
	q := eventSelect + ` WHERE e."deletedAt" IS NULL AND ($1='' OR e."userId"::text=$1)
ORDER BY e."createdAt" DESC`
	data, err := many(p.db, q, userID)
	return data, noRow(err, api.ErrInternal)
}

func (p *Postgres) GetEvent(id string) (map[string]any, api.Error) {
	data, err := one(p.db, eventSelect+` WHERE e.id=$1`, id)
	return data, noRow(err, api.ErrEventNotFound)
}

func (p *Postgres) CreateEvent(userID string, in EventInput) (map[string]any, api.Error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, api.ErrInternal
	}
	defer tx.Rollback()
	var id string
	q := `INSERT INTO event(title,thumbnail,place,"cast","ageLimit",svg,"ticketingStartTime","userId")
VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`
	err = tx.QueryRow(q, in.Title, in.Thumbnail, in.Place, in.Cast,
		in.AgeLimit, in.SVG, nilTime(in.TicketingStartTime), userID).Scan(&id)
	if err != nil {
		return nil, api.ErrInternal
	}
	for _, d := range in.EventDates {
		_, err = tx.Exec(`INSERT INTO event_date(date,"eventId") VALUES($1,$2)`, d.Date, id)
		if err != nil {
			return nil, api.ErrInternal
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, api.ErrInternal
	}
	return p.GetEvent(id)
}

func (p *Postgres) UpdateEvent(id string, in EventInput) (map[string]any, api.Error) {
	q := `UPDATE event SET title=$1,thumbnail=$2,place=$3,"cast"=$4,
"ageLimit"=$5,svg=$6,"ticketingStartTime"=$7,"updatedAt"=now() WHERE id=$8`
	res, err := p.db.Exec(q, in.Title, in.Thumbnail, in.Place, in.Cast,
		in.AgeLimit, in.SVG, nilTime(in.TicketingStartTime), id)
	if err != nil {
		return nil, api.ErrInternal
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return nil, api.ErrEventNotFound
	}
	return p.GetEvent(id)
}

func (p *Postgres) DeleteEvent(id string) api.Error {
	res, err := p.db.Exec(`UPDATE event SET "deletedAt"=now() WHERE id=$1`, id)
	if err != nil {
		return api.ErrInternal
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return api.ErrEventNotFound
	}
	return api.NoError
}

func nilTime(v string) any {
	if v == "" {
		return sql.NullString{}
	}
	return v
}
