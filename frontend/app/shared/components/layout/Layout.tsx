import type { FC, PropsWithChildren } from "react";
import { I18nProps } from "@/app/i18n/props";
import { ChatPanel } from "@/app/shared/components/layout/chatPanel";
import { RoomPanel } from "@/app/shared/components/layout/roomPanel";
import { UserPanel } from "@/app/shared/components/layout/userPanel";
import "./Layout.scss";

type TProps = {
  roomList: any;
} & PropsWithChildren &
  I18nProps;

export const Layout: FC<TProps> = ({ children, i18n, roomList }) => {
  return (
    <div className="Layout">
      <div className="Layout-Box">
        <UserPanel />
        <RoomPanel roomList={roomList} />
        <ChatPanel />
        {/*{children}*/}
      </div>
    </div>
  );
};
