package store

const eventSelect = `SELECT json_build_object(
'id',e.id,'title',e.title,'thumbnail',e.thumbnail,'place',e.place,
'cast',e."cast",'ageLimit',e."ageLimit",'svg',e.svg,
'ticketingStartTime',e."ticketingStartTime",'createdAt',e."createdAt",
'updatedAt',e."updatedAt",'user',json_build_object('id',u.id,
'nickname',u.nickname,'email',u.email,'profileImage',u."profileImage",'role',u.role),
'eventDates',COALESCE((SELECT json_agg(json_build_object('id',d.id,'date',d.date,
'createdAt',d."createdAt",'updatedAt',d."updatedAt")) FROM event_date d
WHERE d."eventId"=e.id AND d."deletedAt" IS NULL),'[]'::json),
'areas',COALESCE((SELECT json_agg(json_build_object('id',a.id,'label',a.label,
'price',a.price,'svg',a.svg,'seats',COALESCE((SELECT json_agg(json_build_object(
'id',s.id,'cx',s.cx,'cy',s.cy,'row',s."row",'number',s.number,
'createdAt',s."createdAt",'updatedAt',s."updatedAt")) FROM seat s
WHERE s."areaId"=a.id AND s."deletedAt" IS NULL),'[]'::json)))
FROM area a WHERE a."eventId"=e.id AND a."deletedAt" IS NULL),'[]'::json)
) FROM event e LEFT JOIN "user" u ON u.id=e."userId"`
