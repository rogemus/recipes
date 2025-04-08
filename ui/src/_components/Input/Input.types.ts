import { Comp } from "@/_models";
import { ChangeEventHandler } from "react";

export type InputProps = {
  label: string;
  type?: HTMLInputElement["type"];
  placeholder?: HTMLInputElement["placeholder"];
  defaultValue?: string;
  onChange?: ChangeEventHandler<HTMLInputElement>;
  id: string;
  name: string;
} & Comp;
