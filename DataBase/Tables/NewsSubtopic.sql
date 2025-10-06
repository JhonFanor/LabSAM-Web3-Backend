CREATE TABLE IF NOT EXISTS "news_subtopic" (
  "news_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("news_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "news_subtopic"
        ADD CONSTRAINT fk_news_subtopic_news
        FOREIGN KEY ("news_id") REFERENCES "news" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "news_subtopic"
        ADD CONSTRAINT fk_news_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;