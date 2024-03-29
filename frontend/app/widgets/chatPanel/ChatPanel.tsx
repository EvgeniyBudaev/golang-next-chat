"use client";

import { type FC, useContext, useEffect, useRef, useState } from "react";
import { type TPagination } from "@/app/api/pagination";
import type { TRoomListItem } from "@/app/api/room/list/types";
import {
  DEFAULT_PAGE,
  DEFAULT_PAGE_LIMIT,
} from "@/app/shared/constants/pagination";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { useSessionNext } from "@/app/shared/hooks";
import {
  ETypeMessage,
  type TMessage,
  WSContent,
} from "@/app/shared/types/message";
import { ChatBody } from "@/app/widgets/chatPanel/chatBody";
import { ChatFooter } from "@/app/widgets/chatPanel/chatFooter";
import { ChatHeader } from "@/app/widgets/chatPanel/chatHeader";
import { getFormattedMessages } from "@/app/widgets/chatPanel/utils";
import "./ChatPanel.scss";

type TProps = {
  isCheckedRoomInProfile: boolean;
  isConnection: boolean;
  onToggleConnection?: (isConnection: boolean) => void;
  roomChecked?: TRoomListItem;
};

export const ChatPanel: FC<TProps> = ({
  isCheckedRoomInProfile,
  isConnection,
  onToggleConnection,
  roomChecked,
}) => {
  const { data: session, status } = useSessionNext();
  const [messageList, setMessageList] = useState<TMessage[]>([]);
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);
  const { conn } = useContext(WebsocketContext);
  const [pagination, setPagination] = useState<TPagination>();
  const [roomId, setRoomId] = useState<number | undefined>(undefined);
  const buttonRef = useRef<HTMLInputElement>(null);

  // console.log("roomId:", roomId);
  // console.log("roomChecked:", roomChecked);
  // useEffect(() => {
  //   setRoomId(roomChecked?.id);
  //   roomId && buttonRef.current && buttonRef.current.click();
  // }, [roomId, roomChecked]);

  useEffect(() => {
    console.log("conn: ", conn);
  }, [conn]);

  useEffect(() => {
    console.log("pagination: ", pagination);
  }, [pagination]);

  const handleChangePagination = (pagination: TPagination) => {
    setPagination(pagination);
  };

  useEffect(() => {
    // if (textarea.current) {
    //   autosize(textarea.current)
    // }
    if (conn === null) {
      // router.push("/");
      return;
    }
    conn.addEventListener("message", (message) => {
      const m: WSContent = JSON.parse(message.data);
      console.log("addEventListener m: ", m);
      if (m.message.isJoined) {
        setMessageList(m.messageListByRoom.content);
      }
      if (m.message.isLeft) {
        setMessageList(m.messageListByRoom.content);
        return;
      }
      if (session?.user.id) {
        session?.user.id === m.message.userId
          ? (m.message.type = ETypeMessage.Self)
          : (m.message.type = ETypeMessage.Recv);
        console.log("Number(m.message.roomId):", Number(m.message.roomId));
        setRoomId(Number(m.message.roomId));
        const formattedMessages = getFormattedMessages(
          m.messageListByRoom.content,
          session?.user.id,
        );
        setMessageList(formattedMessages);
        const { content, ...pagination } = m.messageListByRoom;
        handleChangePagination(pagination);
      }
    });
    conn.onclose = () => {
      console.log("conn.onclose");
      onToggleConnection?.(false);
    };
    conn.onerror = () => {
      onToggleConnection?.(false);
    };
    conn.onopen = () => {
      console.log("conn.onopen");
      onToggleConnection?.(true);
    };
  }, [
    textareaRef,
    conn,
    session?.user.id,
    onToggleConnection,
    messageList,
    roomId,
  ]);

  const sendMessage = () => {
    if (!textareaRef.current?.value) return;
    if (conn === null) {
      // router.push("/");
      return;
    }
    if ("value" in textareaRef.current) {
      const message = {
        content: textareaRef.current.value,
        page: pagination?.page ?? DEFAULT_PAGE,
        limit: pagination?.limit ?? DEFAULT_PAGE_LIMIT,
      };
      const json = JSON.stringify(message);
      conn.send(json);
    }
    if ("value" in textareaRef.current) {
      textareaRef.current.value = "";
    }
  };

  return (
    <div className="ChatPanel">
      <ChatHeader />
      <ChatBody
        messageList={messageList}
        roomId={roomId}
        userId={session?.user.id}
      />
      <ChatFooter
        isCheckedRoomInProfile={isCheckedRoomInProfile}
        isConnection={isConnection}
        onSendMessage={sendMessage}
        ref={textareaRef}
        roomChecked={roomChecked}
      />
    </div>
  );
};
