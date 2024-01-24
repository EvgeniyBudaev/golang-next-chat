"use client";

import clsx from "clsx";
import { type FC, useContext, useEffect, useRef } from "react";
import { type TProfileListItem } from "@/app/api/profile/list/types";
import { useTranslation } from "@/app/i18n/client";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import { Avatar } from "@/app/uikit/components/avatar";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./GlobalSearchResults.scss";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { WEBSOCKET_URL } from "@/app/shared/constants";
import { BufferedWebSocket } from "@/app/shared/utils/bufferedWebSocket";
import { useSessionNext } from "@/app/shared/hooks";
import {
  DEFAULT_PAGE,
  DEFAULT_PAGE_LIMIT,
} from "@/app/shared/constants/pagination";

type TProps = {
  isCheckedRoomInProfile: boolean;
  isConnection: boolean;
  list: TProfileListItem[];
  onItemChecked?: (room: TProfileListItem) => void;
  itemChecked?: TProfileListItem;
};

export const GlobalSearchResults: FC<TProps> = ({
  isCheckedRoomInProfile,
  isConnection,
  list,
  onItemChecked,
  itemChecked,
}) => {
  const { t } = useTranslation("index");
  const buttonRefs = useRef<any>([]);
  const { setConn } = useContext(WebsocketContext);
  const { data: session, status } = useSessionNext();

  useEffect(() => {
    if (isCheckedRoomInProfile && itemChecked) {
      if (buttonRefs.current && buttonRefs.current[itemChecked.id]) {
        buttonRefs.current[itemChecked.id].click();
      }
    }
  }, [itemChecked]);

  const joinRoom = (item: TProfileListItem) => {
    const url =
      `${WEBSOCKET_URL}/room/join?userId=${session?.user.id}` +
      `&username=${session?.user.username}` +
      `&roomTitle=${session?.user.name}` +
      `&receiverId=${item.id}` +
      `&page=${DEFAULT_PAGE}` +
      `&limit=${DEFAULT_PAGE_LIMIT}`;
    const ws = new BufferedWebSocket(url);
    if (ws.OPEN) {
      setConn(ws);
    }
  };

  const handleProfileChecked = (item: TProfileListItem) => {
    onItemChecked?.(item);
    joinRoom(item);
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
                  itemChecked?.id === item.id,
              })}
              key={item.id}
              onClick={() => handleProfileChecked(item)}
            >
              <div className="GlobalSearchResults-ListItem-Left">
                <Avatar
                  className="GlobalSearchResults-Avatar"
                  size={40}
                  user={item.firstName}
                />
                <div className="GlobalSearchResults-ListItem-Right">
                  {item.firstName}&nbsp;{item.lastName}
                </div>
              </div>
              <div>
                <div className="GlobalSearchResults-ListItem-NotViewed">
                  10000
                </div>
              </div>

              {/*{itemChecked && isCheckedRoomInProfile && !isConnection && (*/}
              {/*  <RoomJoinForm*/}
              {/*    button={*/}
              {/*      <button*/}
              {/*        className="GlobalSearchResults-Join"*/}
              {/*        ref={(ref) => (buttonRefs.current[item.id] = ref)}*/}
              {/*        type="submit"*/}
              {/*      >*/}
              {/*        <Typography*/}
              {/*          value={"join to channel"}*/}
              {/*          variant={ETypographyVariant.TextB2Bold}*/}
              {/*        />*/}
              {/*      </button>*/}
              {/*    }*/}
              {/*    room={item}*/}
              {/*  />*/}
              {/*)}*/}
            </div>
          );
        })}
      </div>
    </div>
  );
};
