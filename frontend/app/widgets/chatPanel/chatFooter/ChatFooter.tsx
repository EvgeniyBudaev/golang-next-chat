import { type ForwardedRef, forwardRef } from "react";
import { EFormFields } from "@/app/widgets/chatPanel/chatBody/enums";
import "./ChatFooter.scss";
import { Icon } from "@/app/uikit/components/icon";

type TProps = {
  onSendMessage: () => void;
};

const Component = (props: TProps, ref: ForwardedRef<HTMLTextAreaElement>) => {
  return (
    <div className="ChatFooter">
      <textarea
        className="ChatFooter-WriteField"
        name={EFormFields.Message}
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
