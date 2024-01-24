import { ETypeMessage, TMessage } from "@/app/shared/types/message";
import { TMessageListItemByRoom } from "@/app/api/message/listByRoom";

export const getFormattedMessages = (
  messageList: TMessageListItemByRoom[] | TMessage[],
  userId: string,
) => {
  if (messageList) {
    return messageList.map((message) => {
      return {
        ...message,
        type:
          message.type === ETypeMessage.Sys
            ? ETypeMessage.Sys
            : userId === message.userId
              ? ETypeMessage.Self
              : ETypeMessage.Recv,
      };
    });
  }
  return [];
};
