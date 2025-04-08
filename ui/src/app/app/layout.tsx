import type { FC, PropsWithChildren } from "react";

const CoreLayout: FC<PropsWithChildren> = ({ children }) => {
  return (
    <section style={{ background: "green", padding: "50px" }}>
      {children}
    </section>
  );
};

export default CoreLayout;
