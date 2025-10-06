CREATE TABLE IF NOT EXISTS "investigation_subtopic" (
  "investigation_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("investigation_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "investigation_subtopic"
        ADD CONSTRAINT fk_investigation_subtopic_investigations
        FOREIGN KEY ("investigation_id") REFERENCES "investigations" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "investigation_subtopic"
        ADD CONSTRAINT fk_investigation_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;