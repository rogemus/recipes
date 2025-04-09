import { GetUnitsResponse } from "./actions.types";

const BASE_API_PATH = process.env.API_PATH;

export const getUnits = async (): Promise<GetUnitsResponse> => {
  const API_PATH = `${BASE_API_PATH}/v1/recipes`;
};

export const getIngredients = async () => {
  const API_PATH = `${BASE_API_PATH}/v1/recipes`;
};

export const createIngredient = async () => {
  const API_PATH = `${BASE_API_PATH}/v1/recipes`;
};

export const createRecipe = async () => {
  const API_PATH = `${BASE_API_PATH}/v1/recipes`;
};
