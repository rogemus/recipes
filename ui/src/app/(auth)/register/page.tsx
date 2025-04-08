import { Suspense } from "react";

import RegisterForm from "./_components/RegisterForm";

import type { Metadata } from "next";


export const metadata: Metadata = {
  title: "Register",
};

const Page = () => {
  return (
    <>
      <h1>Register</h1>
      <Suspense>
        <RegisterForm />
      </Suspense>
    </>
  );
};

export default Page;
