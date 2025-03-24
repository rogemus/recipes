import { FC } from "react";
import { RecipeDetailsPageProps } from "./page.types";
import { getRecipeDetails } from "./_lib/actions";

const Page: FC<RecipeDetailsPageProps> = async ({ params }) => {
  const queryParams = await params;
  const id = queryParams.id;

  const { data, error } = await getRecipeDetails(id);
  console.log({ data, error });

  return <h1>Recipe single {id}</h1>;
};

export default Page;
