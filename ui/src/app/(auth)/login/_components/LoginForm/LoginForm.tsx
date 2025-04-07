"use client";

import { useActionState } from "react";
import { login } from "../../_lib/actions";
import { z } from "zod";
import { FormState } from "@/_models/FormState";
import { LoginFormInputs } from "./LoginForm.types";
import FormErrors from "@/_components/FormErrors/FormErrors";

const initialState: FormState<LoginFormInputs> = {
  fieldErrors: new z.ZodError<LoginFormInputs>([]).format(),
  formErrors: [],
};

const LoginForm = () => {
  const [state, formAction, pending] = useActionState(login, initialState);

  return (
    <>
      <form action={formAction} method="POST">
        <div>
          <label>Email</label>
          <input
            type="email"
            id="email"
            placeholder="Email..."
            name="email"
            defaultValue={"tom@example.com"}
          />
          <FormErrors errors={state?.fieldErrors.email?._errors} />
        </div>
        <div>
          <label>Password</label>
          <input
            type="password"
            id="password"
            placeholder="Password..."
            name="password"
            defaultValue="pa55word"
          />
          <FormErrors errors={state?.fieldErrors.password?._errors} />
        </div>
        <button disabled={pending} type="submit">
          Login
        </button>
      </form>
    </>
  );
};

export default LoginForm;
