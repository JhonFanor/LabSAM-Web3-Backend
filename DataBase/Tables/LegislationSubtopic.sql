CREATE TABLE IF NOT EXISTS "legislation_subtopic" (
  "legislation_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("legislation_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "legislation_subtopic"
        ADD CONSTRAINT fk_legislation_subtopic_legislations
        FOREIGN KEY ("legislation_id") REFERENCES "legislations" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "legislation_subtopic"
        ADD CONSTRAINT fk_legislation_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;