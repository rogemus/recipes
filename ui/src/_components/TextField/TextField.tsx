import FieldError from "../FieldError";

import styles from "./TextField.module.css";

import type { FC } from "react";
import type { TextFieldProps } from "./TextField.types";

const TextField: FC<TextFieldProps> = ({
  defaultValue,
  error,
  id,
  label,
  name = "",
  placeholder,
  testId = "Input",
  type = "text",
  required,
  register,
}) => {
  return (
    <div className={styles.wrapper} data-testid={testId}>
      <label
        className={styles.label}
        data-testid={`${testId}-label`}
        htmlFor={id}
      >
        {label}
      </label>
      <input
        data-testid={`${testId}-input`}
        className={styles.input}
        type={type}
        id={id}
        {...(required && { required: true })}
        {...(register && { ...register(name) })}
        {...(placeholder && { placeholder })}
        {...(defaultValue && { defaultValue })}
      />
      <FieldError error={error} testId={`${testId}-error`} />
    </div>
  );
};

export default TextField;
