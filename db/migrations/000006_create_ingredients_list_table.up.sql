CREATE TABLE IF NOT EXISTS
  "ingredients_list" (
    "id" INTEGER NOT NULL UNIQUE,
    "igredient_id" INTEGER NOT NULL,
    "unit_id" INTEGER NOT NULL,
    "amout" NUMERIC NOT NULL DEFAULT 1,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("unit_id") REFERENCES "utils" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("igredient_id") REFERENCES "ingredients" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
  );

