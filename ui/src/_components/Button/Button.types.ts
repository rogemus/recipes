import type { MouseEventHandler } from "react";
import type { Comp } from "@/_models";

export type ButtonProps = {
  disabled?: boolean;
  label: string;
  type?: HTMLButtonElement["type"];
  onClick?: MouseEventHandler<HTMLButtonElement>;
} & Comp;
