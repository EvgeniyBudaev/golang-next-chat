"use client";

import {
  type FC,
  useContext,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { useSessionNext } from "@/app/shared/hooks";
import { type TMessage, WSContent } from "@/app/shared/types/message";
import { ChatBody } from "@/app/widgets/chatPanel/chatBody";
import { ChatFooter } from "@/app/widgets/chatPanel/chatFooter";
import { ChatHeader } from "@/app/widgets/chatPanel/chatHeader";
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
  const [users, setUsers] = useState<Array<{ userId: string }>>([]);
  const [roomId, setRoomId] = useState<number | undefined>(undefined);
  const buttonRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setRoomId(roomChecked?.id);
    roomId && buttonRef.current && buttonRef.current.click();
  }, [roomId, roomChecked]);

  useEffect(() => {
    console.log("conn: ", conn);
  }, [conn]);

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
      // if (m.content == "A new user has joined the room") {
      //   setUsers([...users, { userId: m.userId }]);
      // }
      // if (m.content == "user left the chat") {
      //   const deleteUser = users.filter((user) => user.userId != m.userId);
      //   setUsers([...deleteUser]);
      //   // setMessageList([...messageList, m]);
      //   return;
      // }
      // session?.user?.id === m.userId ? (m.type = "self") : (m.type = "recv");
      setMessageList(m.messageListByRoom);
      // setRoomId(Number(m.roomId));
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
    // messageList,
    conn,
    // users,
    session?.user?.id,
    onToggleConnection,
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
      };
      const json = JSON.stringify(message);
      conn.send(json);
    }
    if ("value" in textareaRef.current) {
      textareaRef.current.value = "";
    }
  };

  const formattedMessages: TMessage[] | undefined = useMemo(() => {
    if (messageList) {
      return messageList.map((message) => {
        return {
          ...message,
          type:
            message.type === "sys"
              ? "sys"
              : session?.user?.id === message.userId
                ? "self"
                : "recv",
        };
      });
    }
    return undefined;
  }, [session?.user?.id, messageList]);

  return (
    <div className="ChatPanel">
      <ChatHeader />
      <ChatBody messageList={formattedMessages} />
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
