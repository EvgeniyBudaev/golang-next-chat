"use client";

import clsx from "clsx";
import { type ForwardedRef, forwardRef, useState } from "react";
import { Icon } from "@/app/uikit/components/icon";
import { EFormFields } from "@/app/widgets/chatPanel/chatBody/enums";
import "./ChatFooter.scss";

type TProps = {
  onSendMessage: () => void;
};

const Component = (props: TProps, ref: ForwardedRef<HTMLTextAreaElement>) => {
  const [isActive, setIsActive] = useState(false);

  const handleBlur = () => {
    setIsActive(false);
  };

  const handleFocus = () => {
    setIsActive(true);
  };

  return (
    <div className="ChatFooter">
      <textarea
        className={clsx("ChatFooter-WriteField", {
          "ChatFooter-WriteField__isActive": isActive,
        })}
        name={EFormFields.Message}
        onBlur={handleBlur}
        onFocus={handleFocus}
        ref={ref}
        placeholder={"Write a message"}
        style={{ resize: "none" }}
      />
      <Icon
        className="ChatFooter-IconSend"
        onClick={props.onSendMessage}
        type="ArrowUp"
      />
    </div>
  );
};

export const ChatFooter = forwardRef(Component);
