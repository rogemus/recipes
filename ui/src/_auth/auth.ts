import NextAuth, { DefaultSession } from "next-auth";
import { emailProvider } from "./providers";
import { authConfig } from "./auth.config";
import { Token as AppToken } from "@/_models";

declare module "next-auth" {
  interface Session {
    user: {
      api_token: AppToken["token"];
      api_token_expiry: AppToken["expiry"];
    } & DefaultSession["user"];
  }

  interface User {
    authentication_token: AppToken;
  }

  interface Token {
    api_token: string;
  }
}

export const { auth, signIn, signOut } = NextAuth({
  ...authConfig,

  providers: [emailProvider],
});
