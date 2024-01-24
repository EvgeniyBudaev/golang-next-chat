import {
  ArrowUpIcon,
  AttentionIcon,
  ExitIcon,
  SearchIcon,
  SpinnerIcon,
} from "@/app/uikit/assets/icons";

export type IconType = "ArrowUp" | "Attention" | "Exit" | "Search" | "Spinner";

export const iconTypes = new Map([
  ["ArrowUp", <ArrowUpIcon key="ArrowUpIcon" />],
  ["Attention", <AttentionIcon key="AttentionIcon" />],
  ["Exit", <ExitIcon key="ExitIcon" />],
  ["Search", <SearchIcon key="SearchIcon" />],
  ["Spinner", <SpinnerIcon key="SpinnerIcon" />],
]);
