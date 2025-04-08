import { expect, test, describe } from "vitest";
import { render } from "@testing-library/react";
import { Input } from "./Input";
import { InputProps } from "./Input.types";

const defaultProps: InputProps = {
  label: "test-label",
  id: "test-id",
  name: "test-name",
  placeholder: "Test placeholder",
};

describe("@component/Input", () => {
  test("renders component correctly", () => {
    const { container } = render(<Input {...defaultProps} />);

    expect(container).toMatchSnapshot();
  });
});
