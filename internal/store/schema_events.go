package store

const schemaEvents = `
CREATE TABLE IF NOT EXISTS event (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 title text NOT NULL,
 thumbnail text,
 place text NOT NULL,
 "cast" text NOT NULL,
 "ageLimit" integer,
 svg text,
 "ticketingStartTime" timestamp,
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 "userId" uuid REFERENCES "user"(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS event_date (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 date timestamp NOT NULL,
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 "eventId" uuid REFERENCES event(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS area (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 label text,
 price integer DEFAULT 0,
 svg text,
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 "eventId" uuid REFERENCES event(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS seat (
 id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
 cx integer NOT NULL,
 cy integer NOT NULL,
 "row" integer NOT NULL,
 number integer NOT NULL,
 "createdAt" timestamp DEFAULT now(),
 "updatedAt" timestamp DEFAULT now(),
 "deletedAt" timestamp,
 "areaId" uuid REFERENCES area(id) ON DELETE CASCADE,
 UNIQUE("areaId", "row", number)
)`
