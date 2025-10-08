CREATE TABLE IF NOT EXISTS "legislations" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "description" text NOT NULL,
  "logo" text,
  "date" date NOT NULL,
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "type_of_law_id" integer NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "legislations"
        ADD CONSTRAINT fk_legislations_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "legislations"
        ADD CONSTRAINT fk_legislations_type_of_law
        FOREIGN KEY ("type_of_law_id") REFERENCES "users" ("type_of_law")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;