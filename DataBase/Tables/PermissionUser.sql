CREATE TABLE IF NOT EXISTS "permission_user" (
  "permission_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  PRIMARY KEY ("permission_id", "user_id")
);

ALTER TABLE "permission_user" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "permission_user" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id")
  ON DELETE CASCADE ON UPDATE CASCADE;
