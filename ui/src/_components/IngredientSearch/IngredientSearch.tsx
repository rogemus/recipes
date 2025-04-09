"use client";

import { useState, type ChangeEvent, type FC } from "react";
import { useDebouncedCallback } from "use-debounce";

import IngredientSearchResults from "../IngredientSearchResults";
import { getIngredientsSearch } from "../../_lib";

import type { IngredientSimple } from "@/_models";
import type { IngredientSearchProps } from "./IngredientSearch.types";

const IngredientSearch: FC<IngredientSearchProps> = () => {
  const [results, setResults] = useState<IngredientSimple[]>([]);

  const fetchIngredients = async (query: string) => {
    try {
      const response = await getIngredientsSearch(query || "");

      if (response?.data?.ingredients) {
        setResults(response.data.ingredients);
      }
    } catch {
      setResults([]);
      console.log("ERROR: cannot fetch recipes");
    }
  };

  const handleChange = useDebouncedCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      const query = e.target.value;

      if (query.length >= 3) {
        fetchIngredients(query);
      }
    },
    300,
  );

  return (
    <div>
      <>IngredientSearch</>
      <div>
        <label>Ingredient</label>
        <input
          name="ingredient"
          id="ingredient"
          onChange={handleChange}
          placeholder="Ingredient ..."
        />
      </div>
      <IngredientSearchResults ingredients={results} />
    </div>
  );
};

export default IngredientSearch;
