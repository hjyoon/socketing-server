package store

const schemaReservations = `
CREATE TABLE IF NOT EXISTS "order" (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 "userId" uuid REFERENCES "user"(id) ON DELETE CASCADE,
 "paymentMethod" text DEFAULT 'socket_pay',
 "canceledAt" timestamp
);
CREATE TABLE IF NOT EXISTS reservation (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 "orderId" uuid REFERENCES "order"(id) ON DELETE CASCADE,
 "seatId" uuid NOT NULL REFERENCES seat(id) ON DELETE CASCADE,
 "eventDateId" uuid NOT NULL REFERENCES event_date(id) ON DELETE CASCADE,
 "canceledAt" timestamp
);
CREATE UNIQUE INDEX IF NOT EXISTS unique_eventdate_seat_canceledat_null
ON reservation("seatId", "eventDateId") WHERE "canceledAt" IS NULL`
