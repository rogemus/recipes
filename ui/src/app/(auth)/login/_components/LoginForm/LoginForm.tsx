"use client";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useActionState } from "react";

import TextField from "@/_components/TextField";

import { login } from "../../_lib/actions";

import { LoginFormSchema } from "./LoginForm.schema";

import type { LoginFormInputs } from "./LoginForm.types";
import type { FormState } from "@/_models/FormState";
import { Button } from "@/_components/Button";

const initialState: FormState<LoginFormInputs> = {
  fieldErrors: new z.ZodError<LoginFormInputs>([]).format(),
  formErrors: [],
};

const LoginForm = () => {
  const {
    register,
    formState: { isValid, errors },
  } = useForm<LoginFormInputs>({
    mode: "all",
    resolver: zodResolver(LoginFormSchema),
  });
  const [, formAction] = useActionState<FormState<LoginFormInputs>, FormData>(
    login,
    initialState,
  );

  return (
    <>
      <form action={formAction} method="POST">
        <div>
          <TextField
            {...register("email")}
            type="email"
            id="email"
            placeholder="Email..."
            label="Email"
            // defaultValue={"tom@example.com"}
            testId="EmailField"
            required
            error={errors?.email?.message as string}
          />
        </div>
        <div>
          <TextField
            {...register("password")}
            label="Password"
            type="password"
            id="password"
            placeholder="Password..."
            // defaultValue="pa55word"
            testId="PasswordField"
            required
            error={errors?.password?.message as string}
          />
        </div>
        <Button
          label="Login"
          type="submit"
          testId="FormSubmit"
          disabled={!isValid}
        />
      </form>
    </>
  );
};

export default LoginForm;
