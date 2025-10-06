CREATE TABLE IF NOT EXISTS "users" (
  "id" serial PRIMARY KEY,
  "email" varchar(255) UNIQUE NOT NULL,
  "avatar" text,
  "password" varchar(255) NOT NULL,
  "email_verified" bool DEFAULT false,
  "email_verification_token" varchar(255),
  "role_id" integer NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id")
  ON DELETE SET NULL ON UPDATE CASCADE;
