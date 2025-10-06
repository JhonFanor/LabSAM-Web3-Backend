CREATE TABLE IF NOT EXISTS "educational_offer_subtopic" (
  "educational_offer_id" integer,
  "subtopic_id" integer,
  PRIMARY KEY ("educational_offer_id", "subtopic_id")
);

ALTER TABLE "educational_offer_subtopic" ADD FOREIGN KEY ("educational_offer_id") REFERENCES "educational_offers" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "educational_offer_subtopic" ADD FOREIGN KEY ("subtopic_id") REFERENCES "subtopics" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
