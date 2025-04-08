import type { User } from "@/_models";

export type SignUpResponse = {
  data?: {
    user: User;
  };
  error?:
    | string
    | {
        email: string;
        password: string;
        name: string;
      };
};
