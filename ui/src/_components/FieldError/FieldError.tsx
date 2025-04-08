import type { FC } from "react";
import type { FieldErrorProps } from "./FieldError.types";

const FieldError: FC<FieldErrorProps> = ({
  error = "",
  testId = "FieldError",
}) => {
  if (error === "") return null;

  return (
    <div data-testid={testId}>
      <p>{error}</p>
    </div>
  );
};

export default FieldError;
