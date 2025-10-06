CREATE TABLE IF NOT EXISTS "bank_of_resume_subtopic" (
  "bank_of_resume_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("bank_of_resume_id", "subtopic_id")
);

ALTER TABLE "bank_of_resume_subtopic" ADD FOREIGN KEY ("bank_of_resume_id") REFERENCES "bank_of_resumes" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "bank_of_resume_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
