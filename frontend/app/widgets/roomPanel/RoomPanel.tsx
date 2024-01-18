"use client";

import isEmpty from "lodash/isEmpty";
import { type ChangeEvent, type FC, useState } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { GlobalSearchResults } from "@/app/entities/globalSearchResults";
import { Search } from "@/app/entities/search";
import { RoomCreateForm } from "@/app/features/room/roomCreateForm";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import "./RoomPanel.scss";
import { Avatar } from "@/app/uikit/components/avatar";
import clsx from "clsx";

type TProps = {
  onRoomChecked?: (room: TRoomListItem) => void;
  roomChecked?: TRoomListItem;
  roomListByProfile: TRoomListItem[];
};

export const RoomPanel: FC<TProps> = ({
  onRoomChecked,
  roomChecked,
  roomListByProfile,
}) => {
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

  const handleRoomChecked = (item: TRoomListItem) => {
    onRoomChecked?.(item);
  };

  return (
    <div className="RoomPanel">
      <Search
        onChangeInputValue={handleChangeSearchInputValue}
        onChangeSearchState={handleChangeSearchState}
      />
      {isSearchActive && (
        <GlobalSearchResults
          list={searchListState}
          onRoomChecked={onRoomChecked}
          roomChecked={roomChecked}
        />
      )}
      {!isSearchActive && (
        <div>
          {/*<RoomCreateForm />*/}
          <div>
            <div className="RoomPanel-List">
              {(roomListByProfile ?? []).map((room: TRoomListItem) => {
                return (
                  <div
                    className={clsx("RoomPanel-ListItem", {
                      ["RoomPanel-ListItem__isChecked"]:
                        roomChecked?.uuid === room.uuid,
                    })}
                    key={room.uuid}
                    onClick={() => handleRoomChecked(room)}
                  >
                    <Avatar
                      className="GlobalSearchResults-Avatar"
                      size={40}
                      user={room.title}
                    />
                    <div>{room.title}</div>
                    {/*<RoomJoinForm room={room} />*/}
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
