CREATE TABLE IF NOT EXISTS "job_board_subtopic" (
  "job_board_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("job_board_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "job_board_subtopic"
        ADD CONSTRAINT fk_job_board_subtopic_jobs_board
        FOREIGN KEY ("job_board_id") REFERENCES "jobs_board" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "job_board_subtopic"
        ADD CONSTRAINT fk_job_board_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;