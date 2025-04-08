"use server";

import { z } from "zod";

import { signIn } from "@/_auth";

import { LoginFormSchema } from "../_components/LoginForm";

import type { LoginFormInputs} from "../_components/LoginForm";
import type { FormState } from "@/_models/FormState";


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
