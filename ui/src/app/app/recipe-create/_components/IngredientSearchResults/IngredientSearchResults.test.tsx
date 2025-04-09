import { expect, test, describe } from "vitest";
import { render } from "@testing-library/react";

import IngredientSearchResults from "./IngredientSearchResults";

import type { IngredientSearchResultsProps } from "./IngredientSearchResults.types";

const defaultProps: IngredientSearchResultsProps = {
  ingredients: [
    {
      id: 5,
      name: "test ingredient",
    },
  ],
};

describe("@component/IngredientSearchResults", () => {
  test("renders component correctly", () => {
    const { container } = render(<IngredientSearchResults {...defaultProps} />);

    expect(container).toMatchSnapshot();
  });
});
