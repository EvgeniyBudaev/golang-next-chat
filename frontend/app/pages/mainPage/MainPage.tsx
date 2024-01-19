"use client";

import { type FC, useMemo, useState } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { UserPanel } from "@/app/widgets/userPanel";
import { RoomPanel } from "@/app/widgets/roomPanel";
import { ChatPanel } from "@/app/widgets/chatPanel";
import "./MainPage.scss";

type TProps = {
  roomListByProfile: TRoomListItem[];
};

export const MainPage: FC<TProps> = ({ roomListByProfile = [] }) => {
  const [roomChecked, setRoomChecked] = useState<TRoomListItem | undefined>();
  const isCheckedRoomInProfile = useMemo(() => {
    return (roomListByProfile ?? []).some(
      (room) => room.id === roomChecked?.id,
    );
  }, [roomChecked, roomListByProfile]);
  const [isConnection, setIsConnection] = useState(false);
  console.log("isConnection: ", isConnection);

  const handleRoomChecked = (item: TRoomListItem) => {
    setRoomChecked(item);
  };

  const handleToggleConnection = (isConnection: boolean) => {
    setIsConnection(isConnection);
  };

  return (
    <div className="MainPage">
      <div className="MainPage-Box">
        <UserPanel />
        <RoomPanel
          isCheckedRoomInProfile={isCheckedRoomInProfile}
          isConnection={isConnection}
          roomChecked={roomChecked}
          roomListByProfile={roomListByProfile}
          onRoomChecked={handleRoomChecked}
        />
        <ChatPanel
          isCheckedRoomInProfile={isCheckedRoomInProfile}
          isConnection={isConnection}
          onToggleConnection={handleToggleConnection}
          roomChecked={roomChecked}
        />
      </div>
    </div>
  );
};
