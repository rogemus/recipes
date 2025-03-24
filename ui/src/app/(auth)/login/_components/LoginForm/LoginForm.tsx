"use client";

import { useActionState } from "react";
import { login } from "../../_lib/actions";

const initialState = {
  error: {
    email: "",
    password: "",
  },
};

const LoginForm = () => {
  const [state, formAction, pending] = useActionState(login, initialState);

  // TODO: handle this case
  if (typeof state?.error === "string") {
    return null;
  }

  return (
    <>
      <form action={formAction} method="POST">
        <div>
          <label>Email</label>
          <input
            required
            type="email"
            id="email"
            placeholder="Email..."
            name="email"
            value={"tom@example.com"}
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
            value="pa55word"
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

export default LoginForm;
