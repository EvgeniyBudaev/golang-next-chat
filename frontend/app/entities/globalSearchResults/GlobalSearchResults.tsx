"use client";

import { type FC } from "react";
import { TUser } from "@/app/api/user/list/types";
import { useTranslation } from "@/app/i18n/client";
import { Avatar } from "@/app/uikit/components/avatar";
import "./GlobalSearchResults.scss";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";

type TProps = {
  userList: TUser[];
};

export const GlobalSearchResults: FC<TProps> = ({ userList }) => {
  const { t } = useTranslation("index");
  console.log("userList: ", userList);

  return (
    <div className="GlobalSearchResults">
      <div className="GlobalSearchResults-Header">
        <Typography
          value={t("globalSearchResults.title")}
          variant={ETypographyVariant.TextB3Regular}
        />
      </div>
      <div className="GlobalSearchResults-List">
        {(userList ?? []).map((user) => (
          <div className="GlobalSearchResults-ListItem" key={user.username}>
            <Avatar
              className="GlobalSearchResults-Avatar"
              size={46}
              user={user.username}
            />
            <div>{user.username}</div>
          </div>
        ))}
      </div>
    </div>
  );
};
