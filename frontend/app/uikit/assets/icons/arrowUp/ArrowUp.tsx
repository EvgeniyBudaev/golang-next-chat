import type { FC } from "react";
import { TIconProps } from "@/app/uikit/assets/icons/types";

export const ArrowUpIcon: FC<TIconProps> = ({
  height = 24,
  width = 24,
  ...props
}) => (
  <svg
    height={height}
    width={width}
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 -960 960 960"
    {...props}
  >
    <path d="M440-160v-487L216-423l-56-57 320-320 320 320-56 57-224-224v487h-80Z" />
  </svg>
);
