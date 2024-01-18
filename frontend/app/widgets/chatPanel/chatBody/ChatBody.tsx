import { type FC } from "react";
import { Message } from "@/app/entities/message";
import { type TMessage } from "@/app/shared/types/message";
import "./ChatBody.scss";

type TProps = {
  messageList?: TMessage[] | undefined;
};

export const ChatBody: FC<TProps> = ({ messageList }) => {
  return (
    <div className="ChatBody">
      {(messageList ?? []).map((message: TMessage, index: number) => (
        <Message key={`${message.content}-${index}`} message={message} />
      ))}
    </div>
  );
};
