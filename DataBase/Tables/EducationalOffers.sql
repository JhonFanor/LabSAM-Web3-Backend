CREATE TABLE IF NOT EXISTS "educational_offers" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "institution" varchar(255) NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "cost" Decimal(10,2),
  "description" text NOT NULL,
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "educational_offers" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE SET NULL ON UPDATE CASCADE;
