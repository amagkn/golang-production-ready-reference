BEGIN;

CREATE TABLE IF NOT EXISTS banana
(
    id         UUID PRIMARY KEY,
    name       TEXT,
    status     TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_banana_name ON apple (name);

COMMIT;
