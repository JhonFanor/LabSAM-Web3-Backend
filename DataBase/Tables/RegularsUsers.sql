CREATE TABLE IF NOT EXISTS "regulars_users" (
  "user_id" integer PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "birth_date" date,
  "location_id" integer,
  "contact_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "regulars_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
  
ALTER TABLE regulars_users
  ADD FOREIGN KEY (location_id) REFERENCES locations(id),
  ADD FOREIGN KEY (contact_id) REFERENCES contacts(id);
