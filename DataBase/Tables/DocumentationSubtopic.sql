CREATE TABLE IF NOT EXISTS "documentation_subtopic" (
  "documentation_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("documentation_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "documentation_subtopic"
        ADD CONSTRAINT fk_documentation_subtopic_documentations
        FOREIGN KEY ("documentation_id") REFERENCES "documentations" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "documentation_subtopic"
        ADD CONSTRAINT fk_documentation_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;