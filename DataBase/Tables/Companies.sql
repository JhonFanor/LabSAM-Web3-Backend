CREATE TABLE IF NOT EXISTS "companies" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "industry" varchar(100) NOT NULL,
  "website" varchar(255),
  "email" varchar(255),
  "is_approved" bool,
  "user_id" integer,
  "localitation_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "companies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE "companies" ADD FOREIGN KEY ("localitation_id") REFERENCES "localitations" ("id")
  ON DELETE SET NULL ON UPDATE CASCADE;
