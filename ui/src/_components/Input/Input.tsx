import type { FC } from "react";
import type { InputProps } from "./Input.types";
import styles from "./Input.module.css";

const Input: FC<InputProps> = ({
  type = "text",
  id,
  placeholder,
  name,
  defaultValue,
  label,
  testId = "Input",
  onChange,
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
        {...(onChange && { onChange })}
        {...(placeholder && { placeholder })}
        {...(defaultValue && { defaultValue })}
      />
    </div>
  );
};

export { Input };
