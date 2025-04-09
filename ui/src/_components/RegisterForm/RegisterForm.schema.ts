import { z } from "zod";

export const RegisterFormSchema = z.object({
  name: z
    .string({ required_error: "Name is required" })
    .min(5, {
      message: "Must be 5 or more characters long",
    })
    .max(20, { message: "Must be 20 or fewer characters long" }),
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
