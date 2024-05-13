-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2024-05-13T07:23:23.173Z

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "property" (
  "id" uuid PRIMARY KEY,
  "type" varchar NOT NULL,
  "price" numeric(7,2) NOT NULL,
  "status" varchar NOT NULL DEFAULT 'available',
  "bedroom" int NOT NULL,
  "bathroom" int NOT NULL,
  "location" varchar NOT NULL,
  "size" varchar NOT NULL,
  "contact" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pictures" (
  "id" uuid PRIMARY KEY,
  "property_id" uuid NOT NULL,
  "img_url" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "property"."type" IS 'rent, sale';

ALTER TABLE "property" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");

ALTER TABLE "pictures" ADD FOREIGN KEY ("property_id") REFERENCES "property" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
