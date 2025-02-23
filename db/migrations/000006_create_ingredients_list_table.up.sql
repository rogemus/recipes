CREATE TABLE IF NOT EXISTS
  "ingredients_list" (
    "id" INTEGER NOT NULL UNIQUE,
    "ingredient_id" INTEGER NOT NULL,
    "unit_id" INTEGER NOT NULL,
    "amount" REAL NOT NULL DEFAULT 1,
    "recipe_id" INTEGER NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("unit_id") REFERENCES "utils" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("ingredient_id") REFERENCES "ingredients" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("recipe_id") REFERENCES "recipies" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
  );

