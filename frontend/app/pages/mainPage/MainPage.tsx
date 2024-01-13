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
  console.log("roomList: ", roomList);

  // useEffect(() => {
  //   if (conn === null) {
  //     router.push("/");
  //     return;
  //   }
  //
  //   const roomId = conn.url.split("/")[5];
  //
  //   async function getUsers() {
  //     try {
  //       const res = await fetch(`${API_URL}/room/room/${roomId}/client/list`, {
  //         method: "GET",
  //         headers: {"Content-Type": "application/json"},
  //       });
  //       const data = await res.json();
  //
  //       setUsers(data);
  //     } catch (e) {
  //       console.error(e);
  //     }
  //   }
  //
  //   getUsers();
  // }, []);

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
