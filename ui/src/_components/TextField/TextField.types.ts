import type { FieldValues } from "react-hook-form";
import type {
  ChangeHandler,
  RefCallBack,
  UseFormRegister,
} from "react-hook-form";
import type { Comp } from "@/_models";

export type TextFieldProps = {
  defaultValue?: string;
  disabled?: boolean;
  error?: string;
  id: string;
  label: string;
  name: string;
  placeholder?: HTMLInputElement["placeholder"];
  required?: boolean;
  type?: HTMLInputElement["type"];
  register?: UseFormRegister<FieldValues>;
  onChange?: ChangeHandler;
  onBlur?: ChangeHandler;
  ref?: RefCallBack;
} & Comp;
