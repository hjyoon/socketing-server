package store

import "github.com/hjyoon/socketing-server/internal/api"

func (p *Postgres) ManagerEvent(userID, eventID, dateID string) (map[string]any, api.Error) {
	data, e := p.GetEvent(eventID)
	if e.Code != 0 {
		return nil, e
	}
	var owner string
	err := p.db.QueryRow(`SELECT "userId" FROM event WHERE id=$1`, eventID).Scan(&owner)
	if err != nil || owner != userID {
		return nil, api.ErrEventNotFound
	}
	var sales int
	q := `SELECT COALESCE(sum(a.price),0) FROM reservation r
JOIN seat s ON s.id=r."seatId" JOIN area a ON a.id=s."areaId"
WHERE a."eventId"=$1 AND r."canceledAt" IS NULL`
	_ = dateID
	_ = p.db.QueryRow(q, eventID).Scan(&sales)
	data["totalSales"] = sales
	return data, api.NoError
}
