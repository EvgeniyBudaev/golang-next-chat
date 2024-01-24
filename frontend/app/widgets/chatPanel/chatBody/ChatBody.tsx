import { type FC } from "react";
import { type TMessage } from "@/app/shared/types/message";
import { InfiniteScrollMessages } from "@/app/widgets/chatPanel/infiniteScrollMessages";
import "./ChatBody.scss";

type TProps = {
  messageList: TMessage[];
  roomId?: number;
  userId?: string;
};

export const ChatBody: FC<TProps> = ({ messageList, roomId, userId }) => {
  return (
    <div className="ChatBody">
      <InfiniteScrollMessages
        initialMessages={messageList}
        roomId={roomId}
        userId={userId}
      />
    </div>
  );
};
