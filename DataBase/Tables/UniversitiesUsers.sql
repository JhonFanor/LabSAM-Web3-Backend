CREATE TABLE IF NOT EXISTS "universities_users" (
  "user_id" integer PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "university_type_id" integer NOT NULL,
  "location_id" integer,
  "contact_id" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "universities_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
  
ALTER TABLE universities_users
  ADD CONSTRAINT fk_universities_users_university_type_id FOREIGN KEY (university_type_id) REFERENCES university_types(id),
  ADD CONSTRAINT fk_universities_users_location_id FOREIGN KEY (location_id) REFERENCES locations(id),
  ADD CONSTRAINT fk_universities_users_contact_id FOREIGN KEY (contact_id) REFERENCES contacts(id);
