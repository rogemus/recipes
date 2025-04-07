"use server";

import { User } from "@/_models";
// import { redirect } from "next/navigation";

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

export async function register(_: unknown, formData: FormData) {
  const email = formData.get("email") as string;
  const password = formData.get("password") as string;
  const name = formData.get("name") as string;

  // TODO validate inputs
  if (email && password && name) {
    const res = await signUp(email, name, password);

    if (res.data?.user.id) {
      // return redirect("/login");
    }

    return res;
  }
}
