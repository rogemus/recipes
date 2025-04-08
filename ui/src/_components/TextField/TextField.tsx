import { type FC } from "react";

import FieldError from "../FieldError";

import styles from "./TextField.module.css";

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
  ref,
  onChange,
  onBlur,
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
        name={name}
        {...(ref && { ref })}
        {...(onChange && { onChange })}
        {...(onBlur && { onBlur })}
        {...(required && { required: true })}
        {...(placeholder && { placeholder })}
        {...(defaultValue && { defaultValue })}
      />
      <FieldError error={error} testId={`${testId}-error`} />
    </div>
  );
};

export default TextField;
