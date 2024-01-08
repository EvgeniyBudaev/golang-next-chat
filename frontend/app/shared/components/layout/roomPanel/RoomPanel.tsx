import { type FC } from "react";
import { Search } from "@/app/entities/search";
import "./RoomPanel.scss";

export const RoomPanel: FC = () => {
  return (
    <div className="RoomPanel">
      <Search />
    </div>
  );
};
