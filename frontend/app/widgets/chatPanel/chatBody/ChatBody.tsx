import { type FC } from "react";
import { Message } from "@/app/entities/message";
import { TMessage } from "@/app/pages/appPage/AppPage";
import "./ChatBody.scss";

type TProps = {
  messageList: TMessage[];
};

export const ChatBody: FC<TProps> = ({ messageList }) => {
  return (
    <div className="ChatBody">
      {messageList
        .slice(0)
        .reverse()
        .map((message: TMessage, index: number) => (
          <Message key={`${message.content}-${index}`} message={message} />
        ))}
    </div>
  );
};
