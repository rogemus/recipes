import Pagination from "./_components/Pagination";
import RecipeItem from "./_components/RecipeItem";
import SearchBar from "./_components/SearchBar";
import { getRecipes } from "./_lib";

import type { RecipesListProps } from "./page.types";

const Page: RecipesListProps = async ({ searchParams }) => {
  const { data, error } = await getRecipes(searchParams);
  console.log({ data, error });

  // TODO: handle error
  // TODO: handle no data
  if (typeof data === "undefined") {
    return <div>No data</div>;
  }

  return (
    <>
      <h1>Recipes List</h1>
      <SearchBar />
      <ul>
        {data.recipes.map((recipe) => (
          <RecipeItem key={`recipe-${recipe.id}`} recipe={recipe} />
        ))}
      </ul>

      <Pagination metadata={data.metadata} />
    </>
  );
};

export default Page;
