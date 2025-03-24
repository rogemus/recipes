import { PropsWithChildren } from "react";

export type SearchParams = Promise<{
  [key: string]: string | string[] | undefined;
}>;

// NOTE: Next.js Page
export interface Page<Params = {}> extends PropsWithChildren {
  params: Promise<Params>;
  searchParams: SearchParams;
}
