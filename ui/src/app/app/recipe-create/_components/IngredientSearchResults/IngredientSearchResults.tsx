import type { FC } from "react";
import type { IngredientSearchResultsProps } from "./IngredientSearchResults.types";

const IngredientSearchResults: FC<IngredientSearchResultsProps> = ({
  ingredients = [],
  testId = "IngredientSearchResults",
}) => {
  if (ingredients.length === 0) return null;
  return (
    <ul data-testid={testId}>
      {ingredients.map(({ id, name }) => (
        <li key={id}>{name}</li>
      ))}
    </ul>
  );
};

export default IngredientSearchResults;
