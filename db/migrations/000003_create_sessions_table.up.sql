CREATE TABLE IF NOT EXISTS
  "sessions" (
    "token" CHAR(43),
    "data" BLOB NOT NULL,
    "expiry" TIMESTAMP (6) NOT NULL,
    "created" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("token")
  );
