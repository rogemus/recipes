import Pagination from "./_components/Pagination/Pagination";
import { getRecipes } from "./_lib/actions";
import { RecipesListProps } from "./page.types";

const Page: RecipesListProps = async ({ searchParams }) => {
  const { data, error } = await getRecipes(searchParams);
  console.log({ data, error });

  // TODO: handle error
  // TODO: handle no data
  if (typeof data === "undefined") {
    return <div> No data </div>;
  }

  return (
    <>
      <h1>Recipes List</h1>
      <Pagination metadata={data.metadata} />
    </>
  );
};

export default Page;
