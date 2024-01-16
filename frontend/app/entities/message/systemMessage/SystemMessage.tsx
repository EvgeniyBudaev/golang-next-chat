import { type FC } from "react";
import type { TMessage } from "@/app/shared/types/message";
import "./SystemMessage.scss";

type TProps = {
  message: TMessage;
};

export const SystemMessage: FC<TProps> = ({ message }) => {
  return (
    <div className="SystemMessage">
      <div className="SystemMessage-Text">{message.content}</div>
    </div>
  );
};
