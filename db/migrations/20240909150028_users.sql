-- migrate:up
CREATE TABLE users (
    "id" BIGSERIAL PRIMARY KEY,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

-- migrate:down
DROP TABLE users;