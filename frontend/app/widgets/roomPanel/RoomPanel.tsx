"use client";

import isEmpty from "lodash/isEmpty";
import { type ChangeEvent, type FC, useState } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { TUser } from "@/app/api/user/list/types";
import { GlobalSearchResults } from "@/app/entities/globalSearchResults";
import { Search } from "@/app/entities/search";
import { RoomCreateForm } from "@/app/features/room/roomCreateForm";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import "./RoomPanel.scss";

type TProps = {
  roomList: TRoomListItem[];
};

export const RoomPanel: FC<TProps> = ({ roomList }) => {
  const [isSearchActive, setIsSearchActive] = useState(false);
  const [userListState, setUserListState] = useState<TUser[]>([]);

  const handleChangeSearchInputValue = (
    event: ChangeEvent<HTMLInputElement>,
  ) => {
    setIsSearchActive(!isEmpty(event.target.value));
  };

  const handleChangeSearchState = (userList: TUser[]) => {
    setUserListState(userList);
  };

  return (
    <div className="RoomPanel">
      <Search
        onChangeInputValue={handleChangeSearchInputValue}
        onChangeSearchState={handleChangeSearchState}
      />
      {isSearchActive && <GlobalSearchResults userList={userListState} />}
      {!isSearchActive && (
        <div>
          {/*<RoomCreateForm />*/}
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
      )}
    </div>
  );
};
