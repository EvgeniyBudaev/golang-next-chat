import { type FC } from "react";
import { Search } from "@/app/entities/search";
import { RoomCreateForm } from "@/app/features/room/roomCreateForm";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import "./RoomPanel.scss";

type TProps = {
  roomList: any;
};

export const RoomPanel: FC<TProps> = ({ roomList }) => {
  return (
    <div className="RoomPanel">
      <Search />
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
