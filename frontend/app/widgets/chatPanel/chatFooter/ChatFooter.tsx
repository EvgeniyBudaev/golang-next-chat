"use client";

import clsx from "clsx";
import { type ForwardedRef, forwardRef, useState } from "react";
import { Icon } from "@/app/uikit/components/icon";
import { EFormFields } from "@/app/widgets/chatPanel/chatBody/enums";
import "./ChatFooter.scss";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { RoomJoinForm } from "@/app/features/room/roomJoinForm";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";

type TProps = {
  isCheckedRoomInProfile: boolean;
  onSendMessage: () => void;
  roomChecked?: TRoomListItem;
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
      <div className="ChatFooter-WriteFieldGroup">
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
        {props?.roomChecked && !props.isCheckedRoomInProfile && (
          <RoomJoinForm
            button={
              <button className="ChatFooter-Join" type="submit">
                <Typography
                  value={"join to channel"}
                  variant={ETypographyVariant.TextB2Bold}
                />
              </button>
            }
            room={props?.roomChecked}
          />
        )}
      </div>
      <Icon
        className="ChatFooter-IconSend"
        onClick={props.onSendMessage}
        type="ArrowUp"
      />
    </div>
  );
};

export const ChatFooter = forwardRef(Component);
