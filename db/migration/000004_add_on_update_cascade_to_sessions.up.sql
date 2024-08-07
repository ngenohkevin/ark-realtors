-- Drop the existing foreign key constraint
ALTER TABLE "sessions" DROP CONSTRAINT sessions_username_fkey;

-- Add the new foreign key constraint with ON UPDATE CASCADE
ALTER TABLE "sessions" ADD CONSTRAINT sessions_username_fkey
FOREIGN KEY ("username") REFERENCES "users" ("username") ON UPDATE CASCADE;
