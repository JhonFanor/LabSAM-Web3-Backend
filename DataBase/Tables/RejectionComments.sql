CREATE TABLE IF NOT EXISTS "rejection_comments" (
  "id" serial PRIMARY KEY,
  "resource_type" varchar(50) NOT NULL,   
  "resource_id" integer NOT NULL,         
  "comment" text NOT NULL,                 
  "created_at" timestamp,
  "updated_at" timestamp
);
