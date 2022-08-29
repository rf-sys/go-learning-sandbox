CREATE TABLE "users"
(
    "id"         bigserial PRIMARY KEY,
    "username"   varchar     NOT NULL,
    "password"   varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- ALTER TABLE "users" ADD CONSTRAINT unique_username_constraint UNIQUE (username)