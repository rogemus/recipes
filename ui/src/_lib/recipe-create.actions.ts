"use server";

import type {
  GetIngredientsSearchResponnse,
  GetUnitsResponse,
} from "./recipe-create.actions.types";

const BASE_API_PATH = process.env.API_PATH;

export const getUnits = async (): Promise<GetUnitsResponse> => {
  const API_PATH = `${BASE_API_PATH}/v1/units`;

  try {
    const response = await fetch(API_PATH);
    const json = await response.json();
    return json as GetUnitsResponse;
  } catch (e) {
    const msg = "Error: Unable to fetch units";
    console.error(msg, e);
    return { error: msg };
  }
};

export const getIngredientsSearch = async (
  query: string,
): Promise<GetIngredientsSearchResponnse> => {
  const API_PATH = `${BASE_API_PATH}/v1/search/ingredients?name=${query}`;

  console.log("\n\n", API_PATH, "\n\n");

  try {
    const response = await fetch(API_PATH);
    const json = await response.json();
    return json as GetIngredientsSearchResponnse;
  } catch (e) {
    const msg = "Error: Unable to fetch ingredients";
    console.error(msg, e);
    return { error: msg };
  }
};

export const createIngredient = async () => {
  const API_PATH = `${BASE_API_PATH}/v1/recipes`;
};

export const createRecipe = async () => {
  const API_PATH = `${BASE_API_PATH}/v1/recipes`;
};
