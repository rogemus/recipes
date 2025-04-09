import type { IngredientSimple, Unit } from "@/_models";

export type GetIngredientsResponnse = {
  data?: {
    ingredients: IngredientSimple[];
  };
  error?: string;
};

export type GetUnitsResponse = {
  data?: {
    units: Unit[];
  };
  error?: string;
};

export type GetIngredientsSearchResponnse = {
  data?: {
    ingredients: IngredientSimple[];
  };
  error?: string;
};

export type CreateIngredientRequest = {};

export type CreateRecipeResponse = {
  data?: {};
  error?: string;
};

export type CreateRecipeRequest = {};
