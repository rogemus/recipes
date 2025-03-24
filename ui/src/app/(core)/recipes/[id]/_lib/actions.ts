"use server";

import { IngredientListItem, Recipe, Response } from "@/app/_models";

const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/v1/recipes`;

type RecipeDetailsResponse = Response<{
  ingredients: IngredientListItem[];
  recipe: Recipe;
}>;

export const getRecipeDetails = async (
  id: string,
): Promise<RecipeDetailsResponse> => {
  const url = `${API_PATH}/${id}`;


  try {
    const response = await fetch(url);
    const json = await response.json();

    if (response.status === 200) {
      return { data: json };
    }

    return { ...json };
  } catch (e) {
    const msg = `Error: Unable to fetch recipe with id: ${id}`;
    console.error(msg, e);
    return { error: msg };
  }
};
