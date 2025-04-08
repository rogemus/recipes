"use client";

import { useActionState } from "react";
import { register } from "../../_lib/actions";
import { RegisterFormInputs } from "./RegisterForm.types";
import { FormState } from "@/_models/FormState";
import FormErrors from "@/_components/FormErrors";
import { z } from "zod";

const initialState: FormState<RegisterFormInputs> = {
  fieldErrors: new z.ZodError<RegisterFormInputs>([]).format(),
  formErrors: [],
};

const RegisterForm = () => {
  const [state, formAction, pending] = useActionState(register, initialState);

  return (
    <>
      <form action={formAction} method="POST">
        <div>
          <label>Name</label>
          <input
            required
            type="text"
            id="name"
            placeholder="Name..."
            name="name"
            defaultValue={"Tom"}
          />
          <FormErrors errors={state?.fieldErrors.name?._errors} />
        </div>
        <div>
          <label>Email</label>
          <input
            required
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
            required
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

export default RegisterForm;
