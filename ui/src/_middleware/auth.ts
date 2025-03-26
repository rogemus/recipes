import { NextRequest, NextResponse } from "next/server";

const authPaths = ["/login", "/register"];
const protectedPaths = ["/dashboard"];

export function authMiddleware(req: NextRequest): NextResponse {
  const token = req.cookies.get("token");

  const pathname = req.nextUrl.pathname;

  if (protectedPaths.includes(pathname) && !token) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  if (authPaths.includes(pathname) && token) {
    return NextResponse.redirect(new URL("/dashboard", req.url));
  }

  return NextResponse.next();
}
