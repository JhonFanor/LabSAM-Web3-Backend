CREATE TABLE IF NOT EXISTS "permission_user" (
  "permission_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  PRIMARY KEY ("permission_id", "user_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "permission_user"
        ADD CONSTRAINT fk_permission_user_permissions
        FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "permission_user"
        ADD CONSTRAINT fk_permission_user_users
        FOREIGN KEY ("user_id") REFERENCES "users" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;