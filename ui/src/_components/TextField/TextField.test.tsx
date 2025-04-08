import { expect, test, describe } from "vitest";
import { render } from "@testing-library/react";

import TextField from "./TextField";

import type { TextFieldProps } from "./TextField.types";

const defaultProps: TextFieldProps = {
  label: "test-label",
  id: "test-id",
  name: "test-name",
  placeholder: "Test placeholder",
  error: "Test error",
};

describe("@component/TextField", () => {
  test("renders component correctly", () => {
    const { container } = render(<TextField {...defaultProps} />);

    expect(container).toMatchSnapshot();
  });
});
