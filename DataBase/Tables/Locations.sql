CREATE TABLE IF NOT EXISTS "locations" (
  "id" serial PRIMARY KEY,
  "country" varchar(100) NOT NULL,
  "city" varchar(100) NOT NULL,
  UNIQUE("country", "city")
);
