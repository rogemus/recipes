import type { Page } from "@/app/_models";

type QueryParams = {
  id: string;
};

export type RecipeDetailsPageProps = Page<QueryParams> & {};
