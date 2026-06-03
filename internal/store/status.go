package store

import "github.com/hjyoon/socketing-server/internal/api"

func (p *Postgres) SeatStatus(eventID, dateID, seatID string) (any, api.Error) {
	q := statusSQL + ` WHERE e.id=$1 AND ($2='' OR ed.id::text=$2)`
	args := []any{eventID, dateID}
	if seatID != "" {
		q += ` AND s.id=$3`
		args = append(args, seatID)
		q += ` GROUP BY s.id,a.id`
		data, err := one(p.db, q, args...)
		return data, noRow(err, api.ErrSeatNotFound)
	}
	q += ` GROUP BY s.id,a.id`
	data, err := many(p.db, q, args...)
	return data, noRow(err, api.ErrInternal)
}

const statusSQL = `SELECT json_build_object('id',s.id,'cx',s.cx,'cy',s.cy,
'row',s."row",'number',s.number,'area',json_build_object('id',a.id,
'label',a.label,'price',a.price),'reservations',COALESCE(json_agg(
json_build_object('id',r.id,'order',json_build_object('id',o.id,'createdAt',
o."createdAt",'updatedAt',o."updatedAt",'user',json_build_object('id',u.id,
'nickname',u.nickname,'email',u.email,'profileImage',u."profileImage",
'role',u.role)))) FILTER (WHERE r.id IS NOT NULL),'[]'::json))
FROM seat s JOIN area a ON a.id=s."areaId" JOIN event e ON e.id=a."eventId"
LEFT JOIN event_date ed ON ed."eventId"=e.id
LEFT JOIN reservation r ON r."seatId"=s.id AND r."eventDateId"=ed.id
AND r."canceledAt" IS NULL
LEFT JOIN "order" o ON o.id=r."orderId" AND o."canceledAt" IS NULL
LEFT JOIN "user" u ON u.id=o."userId"`
