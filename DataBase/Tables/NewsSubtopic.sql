CREATE TABLE IF NOT EXISTS "news_subtopic" (
  "news_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("news_id", "subtopic_id")
);

ALTER TABLE "news_subtopic" ADD FOREIGN KEY ("news_id") REFERENCES "news" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "news_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
