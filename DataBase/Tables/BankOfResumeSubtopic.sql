CREATE TABLE IF NOT EXISTS "bank_of_resume_subtopic" (
  "bank_of_resume_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("bank_of_resume_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "bank_of_resume_subtopic"
        ADD CONSTRAINT fk_bank_of_resume_subtopic_bank_of_resumes
        FOREIGN KEY ("bank_of_resume_id") REFERENCES "bank_of_resumes" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "bank_of_resume_subtopic"
        ADD CONSTRAINT fk_bank_of_resume_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;
