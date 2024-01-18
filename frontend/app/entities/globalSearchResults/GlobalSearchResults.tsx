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

type TProps = {
  list: TRoomListItem[];
};

export const GlobalSearchResults: FC<TProps> = ({ list }) => {
  const { t } = useTranslation("index");
  console.log("roomList: ", list);

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
            <div className="GlobalSearchResults-ListItem" key={item.uuid}>
              <Avatar
                className="GlobalSearchResults-Avatar"
                size={46}
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
