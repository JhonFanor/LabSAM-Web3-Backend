CREATE TABLE IF NOT EXISTS "business_users" (
  "user_id" integer PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "industry" varchar(100),
  "location_id" integer,
  "contact_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "business_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE business_users
  ADD CONSTRAINT fk_business_users_location_id FOREIGN KEY (location_id) REFERENCES locations(id),
  ADD CONSTRAINT fk_business_users_contact_id FOREIGN KEY (contact_id) REFERENCES contacts(id);

