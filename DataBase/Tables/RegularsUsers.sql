CREATE TABLE IF NOT EXISTS "regulars_users" (
  "user_id" integer PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "birth_date" date,
  "location_id" integer,
  "contact_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "regulars_users"
        ADD CONSTRAINT fk_regulars_users_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "regulars_users"
        ADD CONSTRAINT fk_regulars_users_locations
        FOREIGN KEY ("location_id") REFERENCES "locations" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "regulars_users"
        ADD CONSTRAINT fk_regulars_users_contacts
        FOREIGN KEY ("contact_id") REFERENCES "contacts" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;