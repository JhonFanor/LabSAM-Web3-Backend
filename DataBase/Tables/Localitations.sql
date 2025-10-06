CREATE TABLE IF NOT EXISTS "localitations" (
  "id" serial PRIMARY KEY,
  "address" varchar(255) NOT NULL,
  "latitude" decimal(9,6),
  "longitude" decimal(9,6),
  "created_at" timestamp,
  "updated_at" timestamp
);

