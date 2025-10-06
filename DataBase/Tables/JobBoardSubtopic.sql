CREATE TABLE IF NOT EXISTS "job_board_subtopic" (
  "job_board_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("job_board_id", "subtopic_id")
);

ALTER TABLE "job_board_subtopic" ADD FOREIGN KEY ("job_board_id") REFERENCES "jobs_board" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "job_board_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
