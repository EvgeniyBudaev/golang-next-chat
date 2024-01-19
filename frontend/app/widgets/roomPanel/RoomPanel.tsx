"use client";

import clsx from "clsx";
import isEmpty from "lodash/isEmpty";
import { type ChangeEvent, type FC, useState } from "react";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { GlobalSearchResults } from "@/app/entities/globalSearchResults";
import { Search } from "@/app/entities/search";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import { Avatar } from "@/app/uikit/components/avatar";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./RoomPanel.scss";

type TProps = {
  isCheckedRoomInProfile: boolean;
  isConnection: boolean;
  onRoomChecked?: (room: TRoomListItem) => void;
  roomChecked?: TRoomListItem;
  roomListByProfile: TRoomListItem[];
};

export const RoomPanel: FC<TProps> = ({
  isCheckedRoomInProfile,
  isConnection,
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
          isCheckedRoomInProfile={isCheckedRoomInProfile}
          isConnection={isConnection}
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
                        roomChecked?.id === room.id,
                    })}
                    key={room.id}
                    onClick={() => handleRoomChecked(room)}
                  >
                    <Avatar
                      className="GlobalSearchResults-Avatar"
                      size={40}
                      user={room.title}
                    />
                    <div>{room.title}</div>
                    {roomChecked && (
                      <RoomJoinForm
                        button={
                          <button className="RoomPanel-Join" type="submit">
                            <Typography
                              value={"join to channel"}
                              variant={ETypographyVariant.TextB2Bold}
                            />
                          </button>
                        }
                        room={roomChecked}
                      />
                    )}
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
