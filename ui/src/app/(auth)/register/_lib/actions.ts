"use server";

import { User } from "@/_models";
import { FormState, toZodError } from "@/_models/FormState";
import { RegisterFormInputs } from "../_components/RegisterForm/RegisterForm.types";

import { z } from "zod";
const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/v1/users`;

type SignUpResponse = {
  data?: {
    user: User;
  };
  error?:
    | string
    | {
        email: string;
        password: string;
        name: string;
      };
};

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

const RegisterFormSchema = z.object({
  name: z
    .string({ required_error: "Name is required" })
    .min(5, {
      message: "Must be 5 or more characters long",
    })
    .max(20, { message: "Must be 20 or fewer characters long" }),
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
    return { FormErrors: ["Something went wrong"] };
  }
}
