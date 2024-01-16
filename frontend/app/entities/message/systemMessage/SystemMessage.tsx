import { type FC } from "react";
import type { TMessage } from "@/app/shared/types/message";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./SystemMessage.scss";

type TProps = {
  message: TMessage;
};

export const SystemMessage: FC<TProps> = ({ message }) => {
  return (
    <div className="SystemMessage">
      <div className="SystemMessage-Text">
        <Typography
          value={message.content}
          variant={ETypographyVariant.TextB3Regular}
        />
      </div>
    </div>
  );
};
