import { ZodFormattedError, ZodIssue, ZodIssueCode } from "zod";

export type FormState<T> = {
  fieldErrors: ZodFormattedError<T>;
  formErrors: string[];
};

export const toZodError = (data: Record<string, string>): ZodIssue[] => {
  const keys = Object.keys(data);

  return keys.map((key) => {
    return {
      code: ZodIssueCode.custom,
      path: [key],
      message: data[key],
    };
  });
};
