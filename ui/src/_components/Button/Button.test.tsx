import { expect, test, describe, vi } from "vitest";
import { fireEvent, render, screen } from "@testing-library/react";

import { Button } from "./Button";

import type { ButtonProps } from "./Button.types";

const defaultProps: ButtonProps = {
  label: "Test label",
};

describe("@component/Button", () => {
  test("renders component correctly", () => {
    const { container } = render(<Button {...defaultProps} />);

    expect(container).toMatchSnapshot();
  });

  test("renders disabled component correctly", () => {
    const { container } = render(<Button {...defaultProps} disabled />);

    expect(container).toMatchSnapshot();
  });

  test("onClick triggered", () => {
    const mockClick = vi.fn();
    render(<Button {...defaultProps} onClick={mockClick} />);

    const button = screen.getByTestId("Button");
    fireEvent.click(button);

    expect(mockClick).toHaveBeenCalledOnce();
  });
});
