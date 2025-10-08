CREATE TABLE IF NOT EXISTS "currency_type" (
  "id" serial PRIMARY KEY,
  "code" varchar(10) NOT NULL,
  "name" varchar(100) NOT NULL,
  "symbol" varchar(10) NOT NULL
);