import { Token } from "@/_models";
import Credentials from "next-auth/providers/credentials";
import { z } from "zod";

const BASE_API_PATH = process.env.API_PATH;
const API_PATH = `${BASE_API_PATH}/api/v1/tokens/authentication`;

type SignInResponse = {
  data?: {
    authentication_token: Token;
  };
  error?: string | { password: string; email: string };
};

const authentication = async (
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

    const json = (await response.json()) as SignInResponse;

    return json;
  } catch (e) {
    const msg = "Error: Unable to login";
    console.error(msg, e);
    return { error: msg };
  }
};

export const emailProvider = Credentials({
  credentials: {
    email: {},
    password: {},
  },
  authorize: async (credentials) => {
    const { error } = z
      .object({
        email: z.string().email(),
        password: z.string().min(6),
      })
      .safeParse(credentials);

    if (error) {
      return null;
    }

    try {
      const res = await authentication(
        credentials.email as string,
        credentials.password as string,
      );

      if (res.data) {
        const user = {
          name: "Bob",
          email: "bom@tom.com",
          // TODO: TEMP return jwt on serwer
          authentication_token: {
            ...res.data.authentication_token,
          },
        };

        console.log("\n\n", user, "\n\n");

        return user;
      }
    } catch (error) {
      console.log({ error });
      throw new Error("Invalid credentials");
    }

    throw new Error("Invalid credentials");
  },
});
