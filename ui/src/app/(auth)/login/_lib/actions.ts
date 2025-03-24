"use server";

import { redirect } from "next/navigation";
import { Response } from "@/app/_models";
import { cookies } from "next/headers";

const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/v1/tokens/authentication`;

type SignInResponse = Response<
  {
    authentication_token: {
      token: string;
      expiry: string;
    };
  },
  string | { password: string; email: string }
>;

export const signIn = async (
  email: string,
  password: string,
): Promise<SignInResponse> => {
  try {
    const response = await fetch(API_PATH, {
      method: "POST",
      body: JSON.stringify({
        email,
        password,
      }),
    });

    const json = await response.json();
    if (response.status === 201) {
      // TODO: fix types

      return { data: json };
    }

    return { ...json };
  } catch (e) {
    const msg = "Error: Unable to login";
    console.error(msg, e);
    return { error: msg };
  }
};

export const login = async (_: unknown, formData: FormData) => {
  // TODO validate inputs
  const cookieStore = await cookies();
  const email = formData.get("email") as string;
  const password = formData.get("password") as string;

  if (email && password) {
    try {
      const res = await signIn(email, password);
      const token = res.data?.authentication_token.token || "";
      const expiry = res.data?.authentication_token.expiry || "";
      const expires = new Date(expiry);

      cookieStore.set({
        name: "token",
        value: token,
        httpOnly: true,
        secure: true,
        expires: expires,
      });
    } catch (e) {
      console.log(e);
      return e;
    } finally {
      redirect("/recipes");
    }
  }
};
