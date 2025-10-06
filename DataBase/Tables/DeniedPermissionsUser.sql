CREATE TABLE IF NOT EXISTS "denied_permissions_user" (
  "permission_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  PRIMARY KEY ("permission_id", "user_id")
);

ALTER TABLE "denied_permissions_user"
ADD CONSTRAINT "fk_denied_permissions_user_permission"
FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id")
ON DELETE CASCADE
ON UPDATE CASCADE;


ALTER TABLE "denied_permissions_user"
ADD CONSTRAINT "fk_denied_permissions_user_user"
FOREIGN KEY ("user_id") REFERENCES "users" ("id")
ON DELETE CASCADE
ON UPDATE CASCADE;
