import { auth } from "@/_auth";

import type { Metadata } from "next";



export const metadata: Metadata = {
  title: "Dashboard",
};

const Page = async () => {
  const session = await auth();
  console.log("PAge session\n\n", { session }, session);

  return <h1>Dashboard</h1>;
};

export default Page;
