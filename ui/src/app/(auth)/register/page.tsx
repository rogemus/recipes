import { Suspense } from "react";

import RegisterForm from "./_components/RegisterForm";

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
