CREATE TABLE IF NOT EXISTS
  "recipies" (
    "id" INTEGER NOT NULL UNIQUE,
    "title" VARCHAR NOT NULL UNIQUE,
    "description" TEXT NOT NULL,
    "created" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "user_id" INTEGER NOT NULL,
    "igredient_list_id" INTEGER NOT NULL,
    "instructions_list_id" INTEGER NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("igredient_list_id") REFERENCES "ingredients_list" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
  );
