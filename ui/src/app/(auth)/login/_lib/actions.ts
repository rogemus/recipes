"use server";

import { Token } from "@/_models";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/v1/tokens/authentication`;

type SignInResponse = {
  data?: {
    authentication_token: Token;
  };
  error?: string | { password: string; email: string };
};

export const signIn = async (
  email: string,
  password: string,
): Promise<SignInResponse> => {
  const cookieStore = await cookies();
  try {
    const response = await fetch(API_PATH, {
      method: "POST",
      body: JSON.stringify({
        email,
        password,
      }),
    });

    const json = (await response.json()) as SignInResponse;

    if (json?.data) {
      const tokenData = json.data.authentication_token;
      const token = tokenData.token;
      const expiry = tokenData.expiry;
      const expires = new Date(expiry);

      cookieStore.set({
        name: "token",
        value: token,
        httpOnly: true,
        secure: true,
        expires: expires,
      });
    }

    return json;
  } catch (e) {
    const msg = "Error: Unable to login";
    console.error(msg, e);
    return { error: msg };
  }
};

export const login = async (_: unknown, formData: FormData) => {
  const email = formData.get("email") as string;
  const password = formData.get("password") as string;

  // TODO validate inputs
  if (email && password) {
    const res = await signIn(email, password);

    if (res.data?.authentication_token.token) {
      return redirect("/dashboard");
    }

    return res;
  }
};
