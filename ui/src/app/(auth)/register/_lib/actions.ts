"use server";

import { z } from "zod";

import { toZodError } from "@/_models/FormState";

import { RegisterFormSchema } from "../_components/RegisterForm";

import type { SignUpResponse } from "./actions.types";
import type { RegisterFormInputs } from "../_components/RegisterForm/RegisterForm.types";
import type { FormState } from "@/_models/FormState";

const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/v1/users`;

export const signUp = async (
  email: string,
  name: string,
  password: string,
): Promise<SignUpResponse> => {
  try {
    const response = await fetch(API_PATH, {
      method: "POST",
      body: JSON.stringify({
        email,
        name,
        password,
      }),
    });

    const json = (await response.json()) as SignUpResponse;
    return json;
  } catch (e) {
    const msg = "Error: Unable to register user";
    console.error(msg, e);
    return { error: msg };
  }
};

export async function register(_: unknown, formData: FormData) {
  const { error } = RegisterFormSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
    name: formData.get("name") as string,
  });

  if (error) {
    return {
      fieldErrors: error.format(),
      formErrors: [],
    } as FormState<RegisterFormInputs>;
  }

  const email = formData.get("email") as string;
  const password = formData.get("password") as string;
  const name = formData.get("name") as string;
  try {
    const res = await signUp(email, name, password);

    if (res.error) {
      if (typeof res.error === "string") {
        return {
          fieldErrors: new z.ZodError<RegisterFormInputs>([]).format(),
          formErrors: [res.error],
        };
      }

      return {
        fieldErrors: new z.ZodError<RegisterFormInputs>(
          toZodError(res.error),
        ).format(),
        formErrors: [],
      };
    }

    return {
      fieldErrors: new z.ZodError<RegisterFormInputs>([]).format(),
      formErrors: [],
    };
  } catch {
    return {
      fieldErrors: new z.ZodError<RegisterFormInputs>([]).format(),
      formErrors: ["Something went wrong"],
    };
  }
}
