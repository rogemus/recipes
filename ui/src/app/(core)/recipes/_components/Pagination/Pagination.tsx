import type { FC } from "react";
import type { PaginationProps } from "./Pagination.types";
import PaginationLink from "../PaginationLink";
import { getLinksAfter, getLinksBefore } from "./utils";

const Pagination: FC<PaginationProps> = ({ metadata }) => {
  const { last_page, first_page, current_page } = metadata;

  const linksBeforeCurrent = getLinksBefore(first_page, current_page);
  const linksAfterCurrent = getLinksAfter(last_page, current_page);

  return (
    <ul>
      {current_page !== first_page && (
        <li>
          {" "}
          <PaginationLink label="1" pageNumber={1} />
        </li>
      )}

      {linksBeforeCurrent.hasMore && (
        <li>
          <div> ... </div>
        </li>
      )}

      {linksBeforeCurrent.links.map((link) => (
        <li key={`link-before-${link}`}>
          <PaginationLink label={link.toString()} pageNumber={link} />
        </li>
      ))}

      <li style={{ background: "red" }}>
        <PaginationLink
          label={current_page.toString()}
          disabled={true}
          pageNumber={current_page}
        />
      </li>

      {linksAfterCurrent.links.map((link) => (
        <li key={`link-after-${link}`}>
          <PaginationLink label={link.toString()} pageNumber={link} />
        </li>
      ))}

      {linksAfterCurrent.hasMore && (
        <li>
          <div> ...</div>
        </li>
      )}

      {current_page !== last_page && (
        <li>
          <PaginationLink label={last_page.toString()} pageNumber={last_page} />
        </li>
      )}
    </ul>
  );
};

export default Pagination;
