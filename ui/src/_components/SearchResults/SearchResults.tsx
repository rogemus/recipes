import Link from "next/link";

import styles from "./SearchResults.module.css";

import type { FC } from "react";
import type { SearchResultsProps } from "./SearchResults.types";

const SearchResults: FC<SearchResultsProps> = ({ recipes = [] }) => {
  if (recipes.length === 0) {
    return null;
  }

  return (
    <ul className={styles.list} tabIndex={-1}>
      {recipes.map((recipe) => (
        <li key={recipe.id} className={styles.item} tabIndex={-1}>
          <Link
            tabIndex={1}
            className={styles.link}
            href={`/app/recipes/${recipe.id}`}
          >
            {recipe.title}
          </Link>
        </li>
      ))}
    </ul>
  );
};

export default SearchResults;
