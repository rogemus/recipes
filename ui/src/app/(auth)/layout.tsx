import type { FC, PropsWithChildren } from "react";

const AuthLayout: FC<PropsWithChildren> = ({ children }) => {
  return (
    <section style={{ background: "blue", padding: "50px" }}>
      {children}
    </section>
  );
};

export default AuthLayout;
