export type Ingredient = {
  id: number;
  name: string;
  created_at: string;
  version: number;
};

export type IngredientListItem = {
  ingredient_id: number;
  ingredient_name: string;
  unit_id: number;
  unit_name: string;
  recipe_id: number;
  amount: string;
};
