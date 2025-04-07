import type { FC } from "react";
import type { FormErrorsProps } from "./FormErrors.types";

const FormErrors: FC<FormErrorsProps> = ({ errors }) => {
  if (!errors?.length) return null;

  return (
    <div>
      {errors.map((err) => (
        <p key={err}>{err}</p>
      ))}
    </div>
  );
};

export default FormErrors;
