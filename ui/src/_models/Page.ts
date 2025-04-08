import { NextPage } from "next";

export type SearchParams = Promise<{
  [key: string]: string | string[] | undefined;
}>;

// NOTE: Next.js Page
export type Page<Params extends Record<string, string>> = NextPage<{
  params: Promise<Params>;
  searchParams: SearchParams;
}>;
