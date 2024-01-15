import clsx from "clsx";
import type { FC } from "react";
import { type TMessage } from "@/app/shared/types/message";
import { Avatar } from "@/app/uikit/components/avatar";
import "./Message.scss";

type TProps = {
  message: TMessage;
};

export const Message: FC<TProps> = ({ message }) => {
  const isReceiver = message.type === "recv";

  return (
    <div
      className={clsx("Message", {
        Message__isReceiver: isReceiver,
      })}
    >
      <div className="Message-Block">
        {/*{isReceiver && (*/}
        {/*  <Avatar*/}
        {/*    className="Message-Avatar"*/}
        {/*    size={32}*/}
        {/*    user={message.user_id.toString()}*/}
        {/*  />*/}
        {/*)}*/}
        <div className="Message-Content">
          <div className="Message-Text">{message.content}</div>
          {/*<div>{message.img && <img src={message.img} alt="" />}</div>*/}
        </div>
      </div>
    </div>
  );
};
