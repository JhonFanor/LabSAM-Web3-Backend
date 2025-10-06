CREATE TABLE IF NOT EXISTS "company_subtopic" (
  "company_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("company_id", "subtopic_id")
);

ALTER TABLE "company_subtopic" ADD FOREIGN KEY ("company_id") REFERENCES "companies" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "company_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
