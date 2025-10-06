CREATE TABLE IF NOT EXISTS "event_subtopic" (
  "event_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("event_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "event_subtopic"
        ADD CONSTRAINT fk_event_subtopic_events
        FOREIGN KEY ("event_id") REFERENCES "events" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "event_subtopic"
        ADD CONSTRAINT fk_legislation_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;