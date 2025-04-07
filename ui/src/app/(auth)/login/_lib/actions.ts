"use server";

import { signIn } from "@/_auth";
import { FormState } from "@/_models/FormState";
import { z } from "zod";
import { LoginFormInputs } from "../_components/LoginForm/LoginForm.types";

const LoginFormSchema = z.object({
  email: z
    .string({ required_error: "Email is required" })
    .email({ message: "Invalid email address" }),
  // TODO validate special char
  password: z
    .string({
      required_error: "Password is required",
    })
    .min(8, {
      message: "Must be 8 or more characters long",
    })
    .max(32, {
      message: "Must be 5 or fewer characters long",
    }),
});

export const login = async (_: unknown, formData: FormData) => {
  const { error } = LoginFormSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });

  if (error) {
    return {
      fieldErrors: error.format(),
      formErrors: [],
    } as FormState<LoginFormInputs>;
  }

  await signIn("credentials", formData);

  return {
    fieldErrors: new z.ZodError<LoginFormInputs>([]).format(),
    formErrors: [],
  };
};
