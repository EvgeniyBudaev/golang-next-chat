import { type FC } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { Search } from "@/app/entities/search";
import { RoomCreateForm } from "@/app/features/room/roomCreateForm";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import "./RoomPanel.scss";

type TProps = {
  roomList: TRoomListItem[];
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
                <RoomJoinForm room={room} />
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};
