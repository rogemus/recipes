import { getRecipeDetails } from "@/_lib";

import type { RecipeDetailsPageProps } from "./page.types";

const Page: RecipeDetailsPageProps = async ({ params }) => {
  const queryParams = await params;
  const id = queryParams.id;

  const { data, error } = await getRecipeDetails(id);
  console.log({ data, error });

  return <h1>Recipe single {id}</h1>;
};

export default Page;
