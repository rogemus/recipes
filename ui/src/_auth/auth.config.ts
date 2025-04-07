import { NextAuthConfig } from "next-auth";

export const authConfig = {
  pages: {
    signIn: "/login",
  },
  callbacks: {
    authorized({ auth, request: { nextUrl } }) {
      const isLoggedIn = !!auth?.user;
      const isOnApp = nextUrl.pathname.startsWith("/dashboard");

      if (isOnApp) {
        if (isLoggedIn) {
          return true;
        }

        return false;
      } else if (isLoggedIn) {
        return Response.redirect(new URL("/dashboard", nextUrl));
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
  providers: [],
} satisfies NextAuthConfig;
