import { expect, test, describe } from "vitest";
import { render } from "@testing-library/react";

import { SearchBar } from "./SearchBar";

describe("@component/SearchBar", () => {
  test("renders component correctly", () => {
    const { container } = render(<SearchBar />);

    expect(container).toMatchSnapshot();
  });
});
