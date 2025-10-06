CREATE TABLE IF NOT EXISTS "notifications" (
  "id" serial PRIMARY KEY,
  "sender_id" integer REFERENCES "users"("id")
    ON DELETE SET NULL 
    ON UPDATE CASCADE,
  "receiver_id" integer NOT NULL REFERENCES "users"("id")
    ON DELETE CASCADE 
    ON UPDATE CASCADE,
  "message" text NOT NULL,
  "resource_type" varchar(50), 
  "resource_id" integer,      
  "action" varchar(50) NOT NULL, 
  "is_read" boolean DEFAULT false,
  "created_at" timestamp 
);
