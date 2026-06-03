package store

const schemaPayments = `
CREATE TABLE IF NOT EXISTS payment (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 "orderId" uuid REFERENCES "order"(id) ON DELETE CASCADE,
 "paymentAmount" integer DEFAULT 0,
 "paymentMethod" text,
 "paymentStatus" text DEFAULT 'pending',
 "paidAt" timestamp,
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 UNIQUE("orderId", "paymentMethod")
)`
