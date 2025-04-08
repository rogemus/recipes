import { FieldValues } from "react-hook-form";

import type { UseFormRegister } from "react-hook-form";
import type { Comp } from "@/_models";

export type TextFieldProps = {
  defaultValue?: string;
  error?: string;
  id: string;
  label: string;
  name: string;
  placeholder?: HTMLInputElement["placeholder"];
  required?: boolean;
  type?: HTMLInputElement["type"];
  register?: UseFormRegister<Record<string, string>>;
} & Comp;
