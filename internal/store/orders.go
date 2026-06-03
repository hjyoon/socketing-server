package store

import "github.com/hjyoon/socketing-server/internal/api"

func (p *Postgres) ListOrders(userID, eventID string) ([]map[string]any, api.Error) {
	q := orderSQL + ` WHERE o."userId"=$1 AND o."deletedAt" IS NULL
AND ($2='' OR e.id::text=$2) GROUP BY o.id,u.id,ed.id,e.id ORDER BY o."createdAt" DESC`
	data, err := many(p.db, q, userID, eventID)
	return data, noRow(err, api.ErrInternal)
}

func (p *Postgres) GetOrder(orderID, userID string) (map[string]any, api.Error) {
	q := orderSQL + ` WHERE o.id=$1 AND o."userId"=$2 AND o."deletedAt" IS NULL
GROUP BY o.id,u.id,ed.id,e.id`
	data, err := one(p.db, q, orderID, userID)
	if err != nil {
		return map[string]any{}, api.NoError
	}
	return data, api.NoError
}

func (p *Postgres) CancelOrder(orderID, userID string) api.Error {
	tx, err := p.db.Begin()
	if err != nil {
		return api.ErrInternal
	}
	defer tx.Rollback()
	var canceled any
	err = tx.QueryRow(`SELECT "canceledAt" FROM "order" WHERE id=$1 AND "userId"=$2`,
		orderID, userID).Scan(&canceled)
	if err != nil {
		return noRow(err, api.ErrOrderNotFound)
	}
	if canceled != nil {
		return api.ErrCanceledOrder
	}
	seats, qerr := canceledSeats(tx, orderID)
	if qerr != nil {
		return api.ErrInternal
	}
	_, err = tx.Exec(`UPDATE "order" SET "canceledAt"=now() WHERE id=$1`, orderID)
	if err != nil {
		return api.ErrInternal
	}
	if _, err = tx.Exec(`UPDATE reservation SET "canceledAt"=now() WHERE "orderId"=$1`, orderID); err != nil {
		return api.ErrInternal
	}
	if e := noRow(tx.Commit(), api.ErrInternal); e != api.NoError {
		return e
	}
	_ = p.publishCanceledSeats(seats)
	return api.NoError
}
