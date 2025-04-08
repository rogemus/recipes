"use client";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useActionState } from "react";

import TextField from "@/_components/TextField";

import { register as registerUser } from "../../_lib/actions";

import { RegisterFormSchema } from "./RegisterForm.schema";

import type { RegisterFormInputs } from "./RegisterForm.types";
import type { FormState } from "@/_models/FormState";
import { Button } from "@/_components/Button";

const initialState: FormState<RegisterFormInputs> = {
  fieldErrors: new z.ZodError<RegisterFormInputs>([]).format(),
  formErrors: [],
};

const RegisterForm = () => {
  const {
    register,
    formState: { isValid, errors },
  } = useForm<RegisterFormInputs>({
    mode: "all",
    resolver: zodResolver(RegisterFormSchema),
  });
  const [, formAction] = useActionState<
    FormState<RegisterFormInputs>,
    FormData
  >(registerUser, initialState);

  return (
    <>
      <form action={formAction} method="POST">
        <div>
          <label>Name</label>
          <TextField
            {...register("name")}
            // defaultValue={"tom@example.com"}
            testId="UserField"
            required
            id="name"
            placeholder="Name..."
            name="name"
            defaultValue={"Tom"}
            label="Name"
            error={errors?.name?.message as string}
          />
        </div>
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
          disabled={!isValid}
          testId="BtnSubmit"
          label="Register"
          type="submit"
        />
      </form>
    </>
  );
};

export default RegisterForm;
