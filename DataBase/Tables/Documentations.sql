CREATE TABLE IF NOT EXISTS "documentations" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "author" varchar(100) NOT NULL,
  "description" text NOT NULL,
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "documentations"
        ADD CONSTRAINT fk_documentations_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;