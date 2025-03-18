CREATE TABLE IF NOT EXISTS
  ingredients_list (
    id bigserial PRIMARY KEY,
    amount REAL NOT NULL DEFAULT 1,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    recipe_id bigint NOT NULL REFERENCES recipes ON DELETE CASCADE,
    ingredient_id bigint NOT NULL REFERENCES ingredients ON DELETE CASCADE,
    unit_id bigint NOT NULL REFERENCES units ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS ingredients_list_recipe_id_idx ON ingredients_list (recipe_id);
