CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY,
  "content" varchar NOT NULL,
  "is_done" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "done_at" timestamptz NOT NULL
);