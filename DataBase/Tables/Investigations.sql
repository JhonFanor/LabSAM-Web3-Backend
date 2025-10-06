CREATE TABLE IF NOT EXISTS "investigations" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "date" date NOT NULL,
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "investigations"
        ADD CONSTRAINT fk_investigations_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;