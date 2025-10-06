CREATE TABLE IF NOT EXISTS "documentation_subtopic" (
  "documentation_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("documentation_id", "subtopic_id")
);

ALTER TABLE "documentation_subtopic" ADD FOREIGN KEY ("documentation_id") REFERENCES "documentations" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "documentation_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
