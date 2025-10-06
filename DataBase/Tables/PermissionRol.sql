CREATE TABLE IF NOT EXISTS "permission_rol" (
  "permission_id" integer NOT NULL,
  "role_id" integer NOT NULL,
  PRIMARY KEY ("permission_id", "role_id")
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "permission_rol"
        ADD CONSTRAINT fk_permission_rol_permission
        FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "permission_rol"
        ADD CONSTRAINT fk_permission_rol_role
        FOREIGN KEY ("role_id") REFERENCES "roles" ("id")
        ON DELETE CASCADE ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;

