import { fn } from "@storybook/test";

import TextField from "./TextField";

import type { Meta, StoryObj } from "@storybook/react";

const meta: Meta<typeof TextField> = {
  component: TextField,
  title: "Form/TextField",
  args: {
    id: "id",
    name: "Email",
    type: "email",
    label: "Email",
    placeholder: "Email ...",
    onBlur: fn(),
    onChange: fn(),
  },
};

export default meta;

type Story = StoryObj<typeof TextField>;

export const Default: Story = {};
