import type { Recipe, Metadata, RecipeSimple } from "@/_models";

export type RecipesResponse = {
  data?: {
    recipes: Recipe[];
    metadata: Metadata;
  };
  error: string;
};

export type AutocompleteRecipesResponse = {
  data?: {
    recipes: RecipeSimple[];
  };
  error?: string;
};
