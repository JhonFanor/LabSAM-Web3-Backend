CREATE TABLE IF NOT EXISTS "investigation_subtopic" (
  "investigation_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("investigation_id", "subtopic_id")
);

ALTER TABLE "investigation_subtopic" ADD FOREIGN KEY ("investigation_id") REFERENCES "investigations" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "investigation_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
