package store

import "github.com/hjyoon/socketing-server/internal/api"

func (p *Postgres) CreateAreas(eventID string, in AreaBatch) (map[string]any, api.Error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, api.ErrInternal
	}
	defer tx.Rollback()
	var exists int
	_ = tx.QueryRow(`SELECT count(*) FROM event WHERE id=$1`, eventID).Scan(&exists)
	if exists == 0 {
		return nil, api.ErrEventNotFound
	}
	for _, area := range in.Areas {
		var areaID string
		err = tx.QueryRow(`INSERT INTO area(label,price,svg,"eventId")
VALUES($1,$2,$3,$4) RETURNING id`, area.Label, area.Price, area.SVG, eventID).Scan(&areaID)
		if err != nil {
			return nil, api.ErrInternal
		}
		for _, s := range area.Seats {
			_, err = tx.Exec(`INSERT INTO seat(cx,cy,"row",number,"areaId")
VALUES($1,$2,$3,$4,$5)`, s.Cx, s.Cy, s.Row, s.Number, areaID)
			if err != nil {
				return nil, api.ErrDuplicateSeat
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, api.ErrInternal
	}
	return p.GetEvent(eventID)
}

func (p *Postgres) ListSeats(eventID string) ([]map[string]any, api.Error) {
	q := `SELECT json_build_object('id',s.id,'cx',s.cx,'cy',s.cy,'row',s."row",
'number',s.number,'createdAt',s."createdAt",'updatedAt',s."updatedAt",
'area',json_build_object('id',a.id,'label',a.label,'price',a.price,'svg',a.svg))
FROM seat s JOIN area a ON a.id=s."areaId" WHERE a."eventId"=$1`
	data, err := many(p.db, q, eventID)
	return data, noRow(err, api.ErrInternal)
}

func (p *Postgres) GetSeat(eventID, seatID string) (map[string]any, api.Error) {
	q := `SELECT json_build_object('id',s.id,'cx',s.cx,'cy',s.cy,'row',s."row",
'number',s.number,'area',json_build_object('id',a.id,'label',a.label,
'price',a.price,'event',json_build_object('id',e.id,'title',e.title)))
FROM seat s JOIN area a ON a.id=s."areaId" JOIN event e ON e.id=a."eventId"
WHERE e.id=$1 AND s.id=$2`
	data, err := one(p.db, q, eventID, seatID)
	return data, noRow(err, api.ErrSeatNotFound)
}
