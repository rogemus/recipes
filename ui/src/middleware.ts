import NextAuth from "next-auth";
import { authConfig } from "./_auth/auth.config";

export default NextAuth(authConfig).auth;

export const config = {
  // https://nextjs.org/docs/app/building-your-application/routing/middleware#matcher
  matcher: ["/((?!api|v1|_next/static|_next/image|.*\\.png$).*)"],
};
