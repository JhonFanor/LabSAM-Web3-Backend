CREATE TABLE IF NOT EXISTS "events" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "image" text NOT NULL,
  "description" text NOT NULL,
  "link" text NOT NULL,
  "date" date NOT NULL,
  "is_approved" bool,
  "user_id" integer NOT NULL,
  "localitation_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "events"
        ADD CONSTRAINT fk_events_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "events"
        ADD CONSTRAINT fk_events_localitations
        FOREIGN KEY ("localitation_id") REFERENCES "localitations" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;