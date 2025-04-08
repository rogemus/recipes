import { Suspense } from "react";

import LoginForm from "./_components/LoginForm";

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
