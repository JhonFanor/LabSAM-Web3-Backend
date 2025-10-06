CREATE TABLE IF NOT EXISTS "jobs_board" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "company" varchar(255),
  "description" text NOT NULL,
  "type" varchar(50),
  "salary_range" varchar(100),
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "jobs_board" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE SET NULL ON UPDATE CASCADE;
