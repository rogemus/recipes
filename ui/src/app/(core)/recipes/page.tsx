import { getRecipes } from "./_lib/actions";
import { FC } from "react";
import { RecipesListProps } from "./page.types";

const Page: FC<RecipesListProps> = async ({ searchParams }) => {
  const recipes = await getRecipes(searchParams);
  console.log({ recipes });
  return <h1>Recipes List</h1>;
};

export default Page;
