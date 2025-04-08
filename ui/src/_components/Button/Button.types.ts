import { Comp } from "@/_models";
import { MouseEventHandler } from "react";

export type ButtonProps = {
  disabled?: boolean;
  label: string;
  type?: HTMLButtonElement["type"];
  onClick?: MouseEventHandler<HTMLButtonElement>;
} & Comp;
