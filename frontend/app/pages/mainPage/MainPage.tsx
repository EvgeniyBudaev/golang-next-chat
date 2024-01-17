"use client";

import { type FC } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { UserPanel } from "@/app/widgets/userPanel";
import { RoomPanel } from "@/app/widgets/roomPanel";
import { ChatPanel } from "@/app/widgets/chatPanel";
import "./MainPage.scss";

type TProps = {
  roomList: TRoomListItem[];
};

export const MainPage: FC<TProps> = ({ roomList = [] }) => {
  return (
    <div className="MainPage">
      <div className="MainPage-Box">
        <UserPanel />
        <RoomPanel roomList={roomList} />
        <ChatPanel />
      </div>
    </div>
  );
};
