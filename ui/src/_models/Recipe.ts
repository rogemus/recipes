export type Recipe = {
  id: number;
  title: string;
  created: string;
  description: string;
  steps: string[];
  version: number;
  user_id: number;
  user_name: string;
};

export type RecipeSimple = {
  id: number;
  title: string;
};
