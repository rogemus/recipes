"use server";

import { Metadata, Recipe, SearchParams } from "@/_models";

const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/v1/recipes`;

type RecipesResponse = {
  data?: {
    recipes: Recipe[];
    metadata: Metadata;
  };
  error: string;
};

export const getRecipes = async (
  searchParams: SearchParams,
): Promise<RecipesResponse> => {
  const urlParams = new URLSearchParams();

  try {
    const params = await searchParams;

    if (params.title) urlParams.set("title", String(params.title));
    if (params.page) urlParams.set("page", String(params.page));
    if (params.page_size) urlParams.set("page_size", String(params.page_size));
    if (params.sort) urlParams.set("sort", String(params.sort));
  } catch {
    console.error("ERROR: Unable to parse query params. Path:/recipes");
  }

  const url =
    urlParams.size > 0 ? `${API_PATH}?=${urlParams.toString()}` : API_PATH;

  try {
    const response = await fetch(url);
    const json = (await response.json()) as RecipesResponse;
    return json;
  } catch (e) {
    const msg = "Error: Unable to fetch recipes";
    console.error(msg, e);
    return { error: msg };
  }
};
