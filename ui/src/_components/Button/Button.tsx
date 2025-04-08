import type { FC } from "react";
import type { ButtonProps } from "./Button.types";
import styles from "./Button.module.css";

const Button: FC<ButtonProps> = ({
  label,
  testId = "Button",
  disabled = false,
  type = "button",
  onClick,
}) => {
  return (
    <button
      type={type}
      className={styles.button}
      data-testid={testId}
      {...(disabled && { disabled })}
      {...(onClick && { onClick })}
    >
      {label}
    </button>
  );
};

export { Button };
