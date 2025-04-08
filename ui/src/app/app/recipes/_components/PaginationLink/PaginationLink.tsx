import type { FC } from "react";
import type { PaginationLinkProps } from "./PaginationLink.types";
import Link from "next/link";

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
