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
                            "img_url" varchar NOT NULL,
                            "bedroom" int NOT NULL,
                            "bathroom" int NOT NULL,
                            "location" varchar NOT NULL,
                            "size" varchar NOT NULL,
                            "contact" varchar NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "property" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");