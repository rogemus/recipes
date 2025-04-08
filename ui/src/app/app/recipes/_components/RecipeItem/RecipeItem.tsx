import type { FC } from "react";
import type { RecipeItemProps } from "./RecipeItem.types";
import Link from "next/link";

const RecipeItem: FC<RecipeItemProps> = ({ recipe }) => {
  return (
    <li>
      <Link href={`/recipes/${recipe.id}`}> {recipe.title}</Link>
      <p>{recipe.description}</p>
    </li>
  );
};

export default RecipeItem;
