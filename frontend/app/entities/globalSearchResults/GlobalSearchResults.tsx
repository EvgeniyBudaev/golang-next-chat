"use client";

import { type FC } from "react";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { useTranslation } from "@/app/i18n/client";
import { Avatar } from "@/app/uikit/components/avatar";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./GlobalSearchResults.scss";
import clsx from "clsx";

type TProps = {
  list: TRoomListItem[];
  onRoomChecked?: (room: TRoomListItem) => void;
  roomChecked?: TRoomListItem;
};

export const GlobalSearchResults: FC<TProps> = ({
  list,
  onRoomChecked,
  roomChecked,
}) => {
  const { t } = useTranslation("index");

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
                  roomChecked?.uuid === item.uuid,
              })}
              key={item.uuid}
              onClick={() => handleRoomChecked(item)}
            >
              <Avatar
                className="GlobalSearchResults-Avatar"
                size={40}
                user={item.title}
              />
              <div>{item.title}</div>
            </div>
          );
        })}
      </div>
    </div>
  );
};
