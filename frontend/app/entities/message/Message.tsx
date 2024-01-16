import clsx from "clsx";
import type { FC } from "react";
import { type TMessage } from "@/app/shared/types/message";
import { Avatar } from "@/app/uikit/components/avatar";
import "./Message.scss";
import { SystemMessage } from "@/app/entities/message/systemMessage";

type TProps = {
  message: TMessage;
};

export const Message: FC<TProps> = ({ message }) => {
  const isReceiver = message.type === "recv";
  const isSystem = message.type === "sys";

  return (
    <div
      className={clsx("Message", {
        Message__isReceiver: isReceiver,
        Message__isSystem: isSystem,
      })}
    >
      {isSystem && <SystemMessage message={message} />}
      {!isSystem && (
        <div className="Message-Block">
          {isReceiver && (
            <Avatar
              className="Message-Avatar"
              size={32}
              user={message.profile.firstName}
            />
          )}
          <div className="Message-Content">
            {isReceiver && (
              <div className="Message-Username">{`${message.profile.firstName} ${message.profile?.lastName}`}</div>
            )}
            <div className="Message-Text">{message.content}</div>
            {/*<div>{message.img && <img src={message.img} alt="" />}</div>*/}
          </div>
        </div>
      )}
    </div>
  );
};
