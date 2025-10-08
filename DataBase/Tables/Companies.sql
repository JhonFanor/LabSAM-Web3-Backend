CREATE TABLE IF NOT EXISTS "companies" (
  "id" serial PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "logo" text,
  "industry" varchar(100) NOT NULL,
  "website" varchar(255),
  "email" varchar(255),
  "projects" text,
  "is_approved" bool,
  "user_id" integer,
  "localitation_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "companies"
        ADD CONSTRAINT fk_companies_userS
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "companies"
        ADD CONSTRAINT fk_companies_localitations
        FOREIGN KEY ("localitation_id") REFERENCES "localitations" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;