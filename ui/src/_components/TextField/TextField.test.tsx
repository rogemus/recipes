import { expect, test, describe, vi } from "vitest";
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

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

  test("onChange triggered", async () => {
    const mockOnChange = vi.fn();
    const user = userEvent.setup();

    render(<TextField {...defaultProps} onChange={mockOnChange} />);
    const input = screen.getByTestId("Input-input");

    await user.type(input, "Test Content");

    expect(mockOnChange).toHaveBeenCalled();
    expect(input).toHaveValue("Test Content");
  });

  test("onBlur triggered", async () => {
    const mockOnBlur = vi.fn();

    render(<TextField {...defaultProps} onBlur={mockOnBlur} />);
    const input = screen.getByTestId("Input-input");

    input.focus();
    expect(input).toHaveFocus();

    input.blur();
    expect(mockOnBlur).toHaveBeenCalledOnce();
  });
});
