CREATE TABLE IF NOT EXISTS "subtopics" (
  "id" integer PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "topic_id" integer NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "subtopics"
        ADD CONSTRAINT fk_subtopics_topics
        FOREIGN KEY ("topic_id") REFERENCES "topics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;