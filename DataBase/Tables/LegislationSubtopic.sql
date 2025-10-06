CREATE TABLE IF NOT EXISTS "legislation_subtopic" (
  "legislation_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("legislation_id", "subtopic_id")
);

ALTER TABLE "legislation_subtopic" ADD FOREIGN KEY ("legislation_id") REFERENCES "legislations" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "legislation_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
