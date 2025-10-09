CREATE TABLE IF NOT EXISTS "educational_offers" (
  "id" serial PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "institution" varchar(255) NOT NULL,
  "logo" text,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "cost" Decimal(10,2),
  "description" text NOT NULL,
  "link" text NOT NULL,
  "is_approved" bool,
  "user_id" integer,
  "type_education_id" integer NOT NULL,
  "currency_type_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "educational_offers"
        ADD CONSTRAINT fk_educational_offers_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "educational_offers"
        ADD CONSTRAINT fk_educational_offers_type_education
        FOREIGN KEY ("type_education_id") REFERENCES "type_education" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "educational_offers"
        ADD CONSTRAINT fk_educational_offers_currency_type
        FOREIGN KEY ("currency_type_id") REFERENCES "currency_type" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;