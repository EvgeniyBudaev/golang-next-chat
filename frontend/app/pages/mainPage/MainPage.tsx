"use client";

import { type FC, useState } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { UserPanel } from "@/app/widgets/userPanel";
import { RoomPanel } from "@/app/widgets/roomPanel";
import { ChatPanel } from "@/app/widgets/chatPanel";
import "./MainPage.scss";

type TProps = {
  roomList: TRoomListItem[];
};

export const MainPage: FC<TProps> = ({ roomList = [] }) => {
  const [roomChecked, setRoomChecked] = useState<TRoomListItem | undefined>();

  const handleRoomChecked = (item: TRoomListItem) => {
    setRoomChecked(item);
  };

  return (
    <div className="MainPage">
      <div className="MainPage-Box">
        <UserPanel />
        <RoomPanel
          roomChecked={roomChecked}
          roomList={roomList}
          onRoomChecked={handleRoomChecked}
        />
        <ChatPanel roomChecked={roomChecked} />
      </div>
    </div>
  );
};
