CREATE TABLE IF NOT EXISTS "jobs_board" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "logo" text,
  "company" varchar(255),
  "description" text NOT NULL,
  "type" varchar(50),
  "salary_range" varchar(100),
  "link" text NOT NULL,
  "start_date" date NOT NULL,
  "end_date" date,
  "is_approved" bool,
  "user_id" integer,
  "currency_type_id" integer NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "jobs_board"
        ADD CONSTRAINT fk_jobs_board_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "jobs_board"
        ADD CONSTRAINT fk_jobs_board_currency_type
        FOREIGN KEY ("currency_type_id") REFERENCES "currency_type" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;