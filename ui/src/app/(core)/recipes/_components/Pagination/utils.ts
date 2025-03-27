type PaginatinLinks = {
  links: number[];
  hasMore: boolean;
};

export const getLinksBefore = (
  firstPage: number,
  page: number,
): PaginatinLinks => {
  const diff = page - firstPage;

  if (diff <= 1) {
    return { links: [], hasMore: false };
  }

  if (diff <= 2) {
    return { links: [page - 1], hasMore: false };
  }

  if (diff === 3) {
    return { links: [page - 2, page - 1], hasMore: false };
  }

  return { links: [page - 2, page - 1], hasMore: true };
};

export const getLinksAfter = (
  lastPage: number,
  page: number,
): PaginatinLinks => {
  const diff = lastPage - page;

  if (diff <= 1) {
    return { links: [], hasMore: false };
  }

  if (diff <= 2) {
    return { links: [page + 1], hasMore: false };
  }

  if (diff === 3) {
    return { links: [page + 1, page + 2], hasMore: false };
  }

  return { links: [page + 1, page + 2], hasMore: true };
};
