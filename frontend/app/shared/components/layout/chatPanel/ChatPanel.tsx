import { type FC } from "react";
import { ChatBody } from "@/app/shared/components/layout/chatPanel/chatBody";
import { ChatHeader } from "@/app/shared/components/layout/chatPanel/chatHeader";
import "./ChatPanel.scss";

export const ChatPanel: FC = () => {
  return (
    <div className="ChatPanel">
      <ChatHeader />
      <ChatBody />
    </div>
  );
};
