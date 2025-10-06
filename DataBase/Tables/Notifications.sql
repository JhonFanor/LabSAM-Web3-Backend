CREATE TABLE IF NOT EXISTS "notifications" (
  "id" serial PRIMARY KEY,
  "sender_id" integer,
  "receiver_id" integer,
  "message" text NOT NULL,
  "resource_type" varchar(50), 
  "resource_id" integer,      
  "action" varchar(50) NOT NULL, 
  "is_read" boolean DEFAULT false,
  "created_at" timestamp 
);

DO $$
BEGIN
    BEGIN
        ALTER TABLE "notifications"
        ADD CONSTRAINT fk_notifications_users_sender
        FOREIGN KEY ("sender_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;

    BEGIN
        ALTER TABLE "notifications"
        ADD CONSTRAINT fk_notifications_receiver
        FOREIGN KEY ("receiver_id") REFERENCES "users" ("id")
        ON DELETE SET NULL ON UPDATE CASCADE;
    EXCEPTION
        WHEN duplicate_object THEN
    END;
END
$$;