CREATE TABLE IF NOT EXISTS "permissions" (
  "id" integer PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

