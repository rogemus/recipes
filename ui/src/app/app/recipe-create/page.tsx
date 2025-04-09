import { IngredientSearch } from "@/_components";
import { getUnits } from "@/_lib";

const Page = async () => {
  const units = await getUnits();
  console.log("units:", units);

  return (
    <div>
      <h1>Create recipe</h1>
      <IngredientSearch />
    </div>
  );
};

export default Page;
