import { ZodFormattedError } from "zod";

export type FormState<T> = {
  fieldErrors: ZodFormattedError<T>;
  formErrors: string[];
};
