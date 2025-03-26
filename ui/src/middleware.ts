import { NextRequest, NextResponse } from "next/server";
import { authMiddleware } from "./_middleware/auth";

export function middleware(req: NextRequest) {
  let next: NextResponse;
  next = authMiddleware(req);

  return next;
}

export const config = {
  matcher: ["/:path*"],
};
