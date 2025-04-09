import Link from "next/link";

import type { FC } from "react";
import type { PaginationLinkProps } from "./PaginationLink.types";

const PaginationLink: FC<PaginationLinkProps> = ({
  disabled,
  label,
  pageNumber,
}) => {
  if (disabled) {
    return <div>{label}</div>;
  }

  return (
    <Link href={{ pathname: "/recipes", query: { page: pageNumber } }}>
      {label}
    </Link>
  );
};

export default PaginationLink;
