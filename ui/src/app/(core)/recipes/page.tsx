import { getRecipes } from "./_lib/actions";
import { RecipesListProps } from "./page.types";

const Page: RecipesListProps = async ({ searchParams }) => {
  const { data, error } = await getRecipes(searchParams);
  console.log({ data, error });
  return <h1>Recipes List</h1>;
};

export default Page;
