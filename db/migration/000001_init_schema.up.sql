CREATE TABLE "users" (
                         "id" uuid PRIMARY KEY,
                         "username" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "property" (
                            "id" uuid PRIMARY KEY,
                            "type" varchar NOT NULL,
                            "price" numeric(7,2) NOT NULL,
                            "status" varchar NOT NULL DEFAULT 'available',
                            "pictures" jsonb NOT NULL,
                            "bedroom" int NOT NULL,
                            "bathroom" int NOT NULL,
                            "location" varchar NOT NULL,
                            "size" varchar NOT NULL,
                            "contact" varchar NOT NULL,
                            "owner_id" uuid UNIQUE NOT NULL,
                            "agent_id" uuid UNIQUE NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "agent" (
                         "id" uuid PRIMARY KEY,
                         "phone_number" varchar NOT NULL,
                         "user_id" uuid,
                         "national_id" varchar NOT NULL,
                         "kra_pin" varchar NOT NULL
);

CREATE TABLE "owner" (
                         "id" uuid PRIMARY KEY,
                         "phone_number" varchar NOT NULL,
                         "user_id" uuid,
                         "national_id" varchar NOT NULL
);

CREATE TABLE "property_owner" (
                                  "property_id" uuid,
                                  "owner_id" uuid,
                                  PRIMARY KEY ("property_id")
);

CREATE TABLE "property_agent" (
                                  "property_id" uuid,
                                  "agent_id" uuid,
                                  PRIMARY KEY ("property_id")
);

CREATE UNIQUE INDEX "uq_property_owner" ON "property_owner" ("property_id");

CREATE UNIQUE INDEX "uq_property_agent" ON "property_agent" ("property_id");

ALTER TABLE "property" ADD FOREIGN KEY ("owner_id") REFERENCES "owner" ("id");

ALTER TABLE "property" ADD FOREIGN KEY ("agent_id") REFERENCES "agent" ("id");

ALTER TABLE "agent" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "owner" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "property_owner" ADD FOREIGN KEY ("property_id") REFERENCES "property" ("id");

ALTER TABLE "property_owner" ADD FOREIGN KEY ("owner_id") REFERENCES "owner" ("id");

ALTER TABLE "property_agent" ADD FOREIGN KEY ("property_id") REFERENCES "property" ("id");

ALTER TABLE "property_agent" ADD FOREIGN KEY ("agent_id") REFERENCES "agent" ("id");
