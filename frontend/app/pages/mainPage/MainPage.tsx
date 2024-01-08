"use client";

import { type FC } from "react";
import { RoomCreateForm } from "@/app/features/room/roomCreateForm";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import "./MainPage.scss";

type TProps = {
  roomList: any;
};

export const MainPage: FC<TProps> = ({ roomList = [] }) => {
  console.log("roomList: ", roomList);

  return (
    <div className="MainPage">
      <div>
        <RoomCreateForm />
        <div>
          <div>Available Rooms:</div>
          <div>
            {(roomList ?? []).map((room, index) => (
              <div key={index}>
                <div>
                  <div>Комната:</div>
                  <div>{room.name}</div>
                </div>
                <RoomJoinForm />
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};
