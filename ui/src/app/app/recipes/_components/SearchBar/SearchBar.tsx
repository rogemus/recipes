"use client";

import { useState, type ChangeEvent, type FC } from "react";
import { useDebouncedCallback } from "use-debounce";

import { getAutocompleteRecipes } from "../../_lib/actions";
import { SearchResults } from "../SearchResults";

import type { RecipeSimple } from "@/_models";
import type { SearchBarProps } from "./SearchBar.types";

const SearchBar: FC<SearchBarProps> = ({ testId = "SearchBar" }) => {
  const [results, setResults] = useState<RecipeSimple[]>([]);
  const fetchRecipes = async (query: string) => {
    try {
      const response = await getAutocompleteRecipes(query);

      if (response.data) {
        setResults(response.data.recipes);
      }
    } catch (e) {
      setResults([]);
      console.log("ERROR: cannot fetch recipes");
    }
  };

  const handleInputChange = useDebouncedCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      if (e.target.value.length >= 3) {
        fetchRecipes(e.target.value);
      }
    },
    300,
  );

  return (
    <div data-testid={testId}>
      <div>
        <label>Search recipe</label>
        <input
          type="text"
          data-testid={`${testId}-input`}
          onChange={handleInputChange}
          placeholder="Search recipe..."
        />
      </div>
      <SearchResults recipes={results} />
    </div>
  );
};

export { SearchBar };
