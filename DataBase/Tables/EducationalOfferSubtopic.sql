CREATE TABLE IF NOT EXISTS "educational_offer_subtopic" (
  "educational_offer_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("educational_offer_id", "subtopic_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "educational_offer_subtopic"
        ADD CONSTRAINT fk_educational_offer_subtopic_educational_offers
        FOREIGN KEY ("educational_offer_id") REFERENCES "educational_offers" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "educational_offer_subtopic"
        ADD CONSTRAINT fk_educational_offer_subtopic_subtopics
        FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;