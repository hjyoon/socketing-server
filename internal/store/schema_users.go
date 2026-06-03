package store

const schemaUsers = `
CREATE TABLE IF NOT EXISTS "user" (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 nickname text UNIQUE NOT NULL,
 email text UNIQUE NOT NULL,
 "profileImage" text,
 role text DEFAULT 'user',
 password text NOT NULL,
 salt text NOT NULL,
 point integer DEFAULT 500000,
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp
)`
