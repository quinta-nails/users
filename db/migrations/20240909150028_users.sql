-- migrate:up
CREATE TABLE users (
    "id" BIGSERIAL PRIMARY KEY,
    "telegram_id" BIGINT NOT NULL,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR,
    "username" VARCHAR,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    UNIQUE ("telegram_id")
);

-- migrate:down
DROP TABLE users;