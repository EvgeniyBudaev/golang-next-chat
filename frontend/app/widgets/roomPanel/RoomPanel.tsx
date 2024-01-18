"use client";

import isEmpty from "lodash/isEmpty";
import { type ChangeEvent, type FC, useState } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
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
  const [searchListState, setSearchListState] = useState<TRoomListItem[]>([]);

  const handleChangeSearchInputValue = (
    event: ChangeEvent<HTMLInputElement>,
  ) => {
    setIsSearchActive(!isEmpty(event.target.value));
  };

  const handleChangeSearchState = (list: TRoomListItem[]) => {
    setSearchListState(list);
  };

  return (
    <div className="RoomPanel">
      <Search
        onChangeInputValue={handleChangeSearchInputValue}
        onChangeSearchState={handleChangeSearchState}
      />
      {isSearchActive && <GlobalSearchResults list={searchListState} />}
      {!isSearchActive && (
        <div>
          {/*<RoomCreateForm />*/}
          <div>
            <div>
              {(roomList ?? []).map((room, index) => {
                const roomName = `${room.profile.firstName} ${room.profile?.lastName}`;
                return (
                  <div key={index}>
                    <div>
                      <div>Комната:</div>
                      <div>{roomName}</div>
                    </div>
                    <RoomJoinForm room={room} />
                  </div>
                );
              })}
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
