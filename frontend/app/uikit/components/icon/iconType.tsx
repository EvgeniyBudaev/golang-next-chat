import { AttentionIcon, ExitIcon, SearchIcon } from "@/app/uikit/assets/icons";

export type IconType = "Attention" | "Exit" | "Search";

export const iconTypes = new Map([
  ["Attention", <AttentionIcon key="AttentionIcon" />],
  ["Exit", <ExitIcon key="ExitIcon" />],
  ["Search", <SearchIcon key="SearchIcon" />],
]);
