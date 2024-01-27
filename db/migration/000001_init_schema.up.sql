CREATE TYPE Access AS ENUM (
  'PUBLIC',
  'PRIVATE'
);

CREATE TABLE "files" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "access" Access NOT NULL,
  "path" varchar NOT NULL,
  "createdAt" timestamptz NOT NULL DEFAULT (now()),
  "updatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "files" ("name");
