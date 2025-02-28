CREATE TABLE IF NOT EXISTS
  "sessions" (
    "token" TEXT PRIMARY KEY,
    "data" BYTEA NOT NULL,
    "expiry" TIMESTAMPTZ NOT NULL,
    "created" timestamp(0) with time zone NOT NULL DEFAULT NOW()
  );

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
