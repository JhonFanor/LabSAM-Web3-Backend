CREATE TABLE IF NOT EXISTS "bank_of_resumes" (
  "id" serial PRIMARY KEY,
  "photo" text NOT NULL,
  "title" varchar(255) NOT NULL,
  "summary" text NOT NULL,
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "bank_of_resumes"
        ADD CONSTRAINT fk_bank_of_resumes_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;

