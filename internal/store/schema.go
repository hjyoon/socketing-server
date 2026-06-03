package store

var schema = []string{
	`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
	schemaUsers,
	schemaEvents,
	schemaReservations,
	schemaPayments,
	`CREATE INDEX IF NOT EXISTS idx_event_user ON event("userId")`,
	`CREATE INDEX IF NOT EXISTS idx_area_event ON area("eventId")`,
	`CREATE INDEX IF NOT EXISTS idx_seat_area ON seat("areaId")`,
	`CREATE INDEX IF NOT EXISTS idx_reservation_order ON reservation("orderId")`,
}
