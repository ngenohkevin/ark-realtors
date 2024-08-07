-- Drop the foreign key constraint with ON UPDATE CASCADE
ALTER TABLE "sessions" DROP CONSTRAINT sessions_username_fkey;

-- Add the original foreign key constraint without ON UPDATE CASCADE
ALTER TABLE "sessions" ADD CONSTRAINT sessions_username_fkey
FOREIGN KEY ("username") REFERENCES "users" ("username");
