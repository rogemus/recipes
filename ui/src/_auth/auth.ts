import NextAuth, { DefaultSession } from "next-auth";
import { emailProvider } from "./providers";
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
  pages: {
    signIn: "/login",
  },
  callbacks: {
    authorized({ auth, request: { nextUrl } }) {
      const isLoggedIn = !!auth?.user;
      const isProtected = nextUrl.pathname.startsWith("/app");

      if (isProtected) {
        if (isLoggedIn) {
          return true;
        }

        return false;
      } else if (isLoggedIn) {
        return Response.redirect(new URL("/app/dashboard", nextUrl));
      }

      return true;
    },
    jwt({ token, user }) {
      if (user) {
        token.api_token = user.authentication_token.token;
      }

      return token;
    },
    session({ session, token }) {
      session.user.api_token = token.api_token as string;
      return session;
    },
  },
  providers: [emailProvider],
});
