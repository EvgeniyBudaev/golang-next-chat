"use client";

import { type FC, useEffect, useRef } from "react";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { useTranslation } from "@/app/i18n/client";
import { Avatar } from "@/app/uikit/components/avatar";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./GlobalSearchResults.scss";
import clsx from "clsx";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";

type TProps = {
  isCheckedRoomInProfile: boolean;
  isConnection: boolean;
  list: TRoomListItem[];
  onRoomChecked?: (room: TRoomListItem) => void;
  roomChecked?: TRoomListItem;
};

export const GlobalSearchResults: FC<TProps> = ({
  isCheckedRoomInProfile,
  isConnection,
  list,
  onRoomChecked,
  roomChecked,
}) => {
  const { t } = useTranslation("index");
  const buttonRefs = useRef<any>([]);

  useEffect(() => {
    if (isCheckedRoomInProfile && roomChecked) {
      if (buttonRefs.current && buttonRefs.current[roomChecked.id]) {
        buttonRefs.current[roomChecked.id].click();
      }
    }
  }, [roomChecked]);

  const handleRoomChecked = (item: TRoomListItem) => {
    onRoomChecked?.(item);
  };

  return (
    <div className="GlobalSearchResults">
      <div className="GlobalSearchResults-Header">
        <Typography
          value={t("globalSearchResults.title")}
          variant={ETypographyVariant.TextB3Regular}
        />
      </div>
      <div className="GlobalSearchResults-List">
        {(list ?? []).map((item) => {
          return (
            <div
              className={clsx("GlobalSearchResults-ListItem", {
                ["GlobalSearchResults-ListItem__isChecked"]:
                  roomChecked?.id === item.id,
              })}
              key={item.id}
              onClick={() => handleRoomChecked(item)}
            >
              <Avatar
                className="GlobalSearchResults-Avatar"
                size={40}
                user={item.title}
              />
              <div>{item.title}</div>
              {roomChecked && isCheckedRoomInProfile && !isConnection && (
                <RoomJoinForm
                  button={
                    <button
                      className="GlobalSearchResults-Join"
                      ref={(ref) => (buttonRefs.current[item.id] = ref)}
                      type="submit"
                    >
                      <Typography
                        value={"join to channel"}
                        variant={ETypographyVariant.TextB2Bold}
                      />
                    </button>
                  }
                  room={item}
                />
              )}
            </div>
          );
        })}
      </div>
    </div>
  );
};
