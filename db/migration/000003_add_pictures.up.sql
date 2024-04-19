
CREATE TABLE "pictures" (
                            "id" uuid PRIMARY KEY,
                            "property_id" uuid NOT NULL,
                            "img_url" varchar NOT NULL,
                            "description" varchar NOT NULL
);

ALTER TABLE "pictures" ADD FOREIGN KEY ("property_id") REFERENCES "property" ("id");