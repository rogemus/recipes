CREATE TABLE IF NOT EXISTS
  "recipes" (
    "id" INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "title" VARCHAR NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    "created" timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    "user_id" INTEGER NOT NULL
  );
