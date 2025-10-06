CREATE TABLE IF NOT EXISTS "news" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "image" text NOT NULL,
  "description" text NOT NULL,
  "link" text,
  "date" date NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "news" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE SET NULL ON UPDATE CASCADE;
