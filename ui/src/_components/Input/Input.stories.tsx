import { Input } from "./Input";

import type { Meta, StoryObj } from "@storybook/react";

import { fn } from "@storybook/test";

const meta: Meta<typeof Input> = {
  component: Input,
  title: "Input",
  args: {
    onChange: fn,
    id: "id",
    name: "Email",
    type: "email",
    label: "Email",
    placeholder: "Email ...",
  },
};

export default meta;

type Story = StoryObj<typeof Input>;

export const Default: Story = {};
