import { z } from "zod";

export const LoginFormSchema = z.object({
  email: z
    .string({ required_error: "Email is required" })
    .email({ message: "Invalid email address" }),
  // TODO validate special char
  password: z
    .string({
      required_error: "Password is required",
    })
    .min(8, {
      message: "Must be 8 or more characters long",
    })
    .max(32, {
      message: "Must be 5 or fewer characters long",
    }),
});
