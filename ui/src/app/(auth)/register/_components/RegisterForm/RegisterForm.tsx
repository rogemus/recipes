"use client";

import { useActionState } from "react";
import { register } from "../../_lib/actions";

const initialState = {
  error: {
    email: "",
    name: "",
    password: "",
  },
};

const RegisterForm = () => {
  const [state, formAction, pending] = useActionState(register, initialState);

  // TODO: handle this case
  if (typeof state?.error === "string") {
    return null;
  }

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
          {state?.error?.email && <p>{state?.error?.email}</p>}
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
          {state?.error?.email && <p>{state?.error?.email}</p>}
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

          {state?.error?.password && <p>{state?.error?.password}</p>}
        </div>
        <button disabled={pending} type="submit">
          Login
        </button>
      </form>
    </>
  );
};

export default RegisterForm;
