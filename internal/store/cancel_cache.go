package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const reservationBroadcast = "socketing:reservation:broadcast"

type canceledSeat struct {
	EventID     string
	EventDateID string
	AreaID      string
	SeatID      string
}

func canceledSeats(tx *sql.Tx, orderID string) ([]canceledSeat, error) {
	rows, err := tx.Query(`SELECT e.id,r."eventDateId",s."areaId",s.id
FROM reservation r JOIN seat s ON s.id=r."seatId"
JOIN area a ON a.id=s."areaId" JOIN event e ON e.id=a."eventId"
WHERE r."orderId"=$1 AND r."canceledAt" IS NULL`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []canceledSeat
	for rows.Next() {
		var seat canceledSeat
		if err := rows.Scan(&seat.EventID, &seat.EventDateID, &seat.AreaID, &seat.SeatID); err != nil {
			return nil, err
		}
		out = append(out, seat)
	}
	return out, rows.Err()
}

func (p *Postgres) publishCanceledSeats(seats []canceledSeat) error {
	if p.cache == nil || len(seats) == 0 {
		return nil
	}
	ctx := context.Background()
	groups := map[string][]map[string]any{}
	now := time.Now().Format(time.RFC3339Nano)
	for _, seat := range seats {
		area := seat.EventID + "_" + seat.EventDateID + "_" + seat.AreaID
		update := canceledSeatUpdate(seat.SeatID, now)
		if merged, ok := mergeSeatJSON(ctx, p, area, seat.SeatID, update); ok {
			_ = p.cache.HSet(ctx, "seats:"+area, seat.SeatID, merged).Err()
		}
		_ = p.cache.Del(ctx, "timer:"+area+":"+seat.SeatID).Err()
		groups[area] = append(groups[area], update)
	}
	for area, payload := range groups {
		raw, err := json.Marshal(map[string]any{
			"room": area, "type": "seatsSelected", "payload": payload,
		})
		if err != nil {
			return err
		}
		if err := p.cache.Publish(ctx, reservationBroadcast, raw).Err(); err != nil {
			return err
		}
	}
	return nil
}

func canceledSeatUpdate(id, now string) map[string]any {
	return map[string]any{
		"seatId": id, "selectedBy": nil, "updatedAt": now,
		"expirationTime": nil, "reservedUserId": nil,
	}
}
