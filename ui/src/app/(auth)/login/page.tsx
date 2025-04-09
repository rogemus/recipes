import { Suspense } from "react";

import LoginForm from "./_components/LoginForm";

import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Login",
};

const Page = () => {
  return (
    <>
      <h1>Login</h1>
      <Suspense>
        <LoginForm />
      </Suspense>
    </>
  );
};

export default Page;
