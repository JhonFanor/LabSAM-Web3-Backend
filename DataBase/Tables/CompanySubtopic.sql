CREATE TABLE IF NOT EXISTS "company_subtopic" (
  "company_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("company_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "company_subtopic"
        ADD CONSTRAINT fk_company_subtopic_companies
        FOREIGN KEY ("company_id") REFERENCES "companies" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "company_subtopic"
        ADD CONSTRAINT fk_company_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;