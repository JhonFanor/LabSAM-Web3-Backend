CREATE TABLE IF NOT EXISTS "event_subtopic" (
  "event_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("event_id", "subtopic_id")
);

ALTER TABLE "event_subtopic" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "event_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
