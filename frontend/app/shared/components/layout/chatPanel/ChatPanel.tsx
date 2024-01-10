"use client";

import { type FC, useContext, useEffect, useRef, useState } from "react";
import { Message } from "@/app/pages/appPage/AppPage";
import { ChatBody } from "@/app/shared/components/layout/chatPanel/chatBody";
import { ChatFooter } from "@/app/shared/components/layout/chatPanel/chatFooter";
import { ChatHeader } from "@/app/shared/components/layout/chatPanel/chatHeader";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import "./ChatPanel.scss";
import { useSessionNext } from "@/app/shared/hooks";

export const ChatPanel: FC = () => {
  const { data: session, status } = useSessionNext();
  console.log("session: ", session);
  const [messageList, setMessageList] = useState<Message[]>([]);
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);
  const { conn } = useContext(WebsocketContext);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  console.log("messageList: ", messageList);

  useEffect(() => {
    // if (textarea.current) {
    //   autosize(textarea.current)
    // }

    if (conn === null) {
      // router.push("/");
      return;
    }

    conn.onmessage = (message) => {
      console.log("conn.onmessage: ", message);
      const m: Message = JSON.parse(message.data);
      if (m.content == "A new user has joined the room") {
        setUsers([...users, { username: m.username }]);
      }

      if (m.content == "user left the chat") {
        const deleteUser = users.filter((user) => user.username != m.username);
        setUsers([...deleteUser]);
        setMessageList([...messageList, m]);
        return;
      }

      session?.user?.name == m.username ? (m.type = "self") : (m.type = "recv");
      setMessageList([...messageList, m]);
    };

    conn.onclose = () => {};
    conn.onerror = () => {};
    conn.onopen = () => {};
  }, [textareaRef, messageList, conn, users]);

  const sendMessage = () => {
    if (!textareaRef.current?.value) return;
    if (conn === null) {
      // router.push("/");
      return;
    }

    if ("value" in textareaRef.current) {
      conn.send(textareaRef.current.value);
    }
    if ("value" in textareaRef.current) {
      textareaRef.current.value = "";
    }
  };

  return (
    <div className="ChatPanel">
      <ChatHeader />
      <ChatBody messageList={messageList} />
      <ChatFooter onSendMessage={sendMessage} />
    </div>
  );
};
