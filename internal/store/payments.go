package store

import (
	"github.com/lib/pq"

	"github.com/hjyoon/socketing-server/internal/api"
)

func (p *Postgres) CreatePayment(userID string, in PaymentInput) (map[string]any, api.Error) {
	if in.PaymentMethod == "" {
		in.PaymentMethod = "socket_pay"
	}
	tx, err := p.db.Begin()
	if err != nil {
		return nil, api.ErrInternal
	}
	defer tx.Rollback()
	if !exists(tx, `SELECT count(*) FROM event_date WHERE id=$1`, in.EventDateID) {
		return nil, api.ErrEventDate
	}
	var seatCount int
	err = tx.QueryRow(`SELECT count(*) FROM seat WHERE id=ANY($1)`, pq.Array(in.SeatIDs)).
		Scan(&seatCount)
	if err != nil || seatCount != len(in.SeatIDs) {
		return nil, api.ErrSeatNotFound
	}
	if reserved(tx, in.EventDateID, in.SeatIDs) {
		return nil, api.ErrExistingOrder
	}
	var amount int
	err = tx.QueryRow(`SELECT COALESCE(sum(a.price),0) FROM seat s
JOIN area a ON a.id=s."areaId" WHERE s.id=ANY($1)`, pq.Array(in.SeatIDs)).Scan(&amount)
	if err != nil {
		return nil, api.ErrInternal
	}
	if !payable(tx, userID, amount) {
		return nil, api.ErrBalance
	}
	var orderID string
	err = tx.QueryRow(`INSERT INTO "order"("userId","paymentMethod") VALUES($1,$2) RETURNING id`,
		userID, in.PaymentMethod).Scan(&orderID)
	if err != nil {
		return nil, api.ErrInternal
	}
	for _, id := range in.SeatIDs {
		_, err = tx.Exec(`INSERT INTO reservation("orderId","seatId","eventDateId")
VALUES($1,$2,$3)`, orderID, id, in.EventDateID)
		if err != nil {
			return nil, api.ErrReservation
		}
	}
	_, err = tx.Exec(`UPDATE "user" SET point=point-$1 WHERE id=$2`, amount, userID)
	if err != nil || tx.Commit() != nil {
		return nil, api.ErrInternal
	}
	return p.GetOrder(orderID, userID)
}

func exists(q queryer, sql string, args ...any) bool {
	var n int
	return q.QueryRow(sql, args...).Scan(&n) == nil && n > 0
}

func reserved(q queryer, dateID string, seats []string) bool {
	return exists(q, `SELECT count(*) FROM reservation r JOIN "order" o ON o.id=r."orderId"
WHERE r."eventDateId"=$1 AND r."seatId"=ANY($2) AND r."canceledAt" IS NULL
AND o."deletedAt" IS NULL`, dateID, pq.Array(seats))
}

func payable(q queryer, userID string, amount int) bool {
	var point int
	return q.QueryRow(`SELECT point FROM "user" WHERE id=$1`, userID).Scan(&point) == nil &&
		point >= amount
}
