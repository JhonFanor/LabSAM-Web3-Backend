CREATE TABLE IF NOT EXISTS "denied_permissions_user" (
  "permission_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  PRIMARY KEY ("permission_id", "user_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "denied_permissions_user"
        ADD CONSTRAINT fk_denied_permissions_user_permissions
        FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "denied_permissions_user"
        ADD CONSTRAINT fk_denied_permissions_user_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;