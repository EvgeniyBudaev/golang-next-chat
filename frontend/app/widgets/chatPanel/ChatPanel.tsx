"use client";

import { type FC, useContext, useEffect, useRef, useState } from "react";
import { type TMessage } from "@/app/shared/types/message";
import { ChatBody } from "@/app/widgets/chatPanel/chatBody";
import { ChatFooter } from "@/app/widgets/chatPanel/chatFooter";
import { ChatHeader } from "@/app/widgets/chatPanel/chatHeader";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { useSessionNext } from "@/app/shared/hooks";
import "./ChatPanel.scss";

export const ChatPanel: FC = () => {
  const { data: session, status } = useSessionNext();
  console.log("session: ", session);
  const [messageList, setMessageList] = useState<TMessage[]>([]);
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
      const m: TMessage = JSON.parse(message.data);
      if (m.content == "A new user has joined the room") {
        setUsers([...users, { username: m.username }]);
      }

      if (m.content == "user left the chat") {
        const deleteUser = users.filter((user) => user.username != m.username);
        setUsers([...deleteUser]);
        setMessageList([...messageList, m]);
        return;
      }

      console.log("session?.user?.username: ", session?.user?.username);
      console.log("m.username: ", m.username);
      session?.user?.username === m.username
        ? (m.type = "self")
        : (m.type = "recv");
      setMessageList([...messageList, m]);
    };

    conn.onclose = () => {};
    conn.onerror = () => {};
    conn.onopen = () => {};
  }, [textareaRef, messageList, conn, users]);

  const sendMessage = () => {
    console.log("sendMessage");
    console.log("textareaRef", textareaRef.current?.value);
    if (!textareaRef.current?.value) return;
    if (conn === null) {
      console.log("conn === null", conn);
      // router.push("/");
      return;
    }

    if ("value" in textareaRef.current) {
      console.log("textareaRef_1", textareaRef.current.value);
      conn.send(textareaRef.current.value);
    }
    if ("value" in textareaRef.current) {
      console.log("textareaRef_2", textareaRef.current.value);
      textareaRef.current.value = "";
    }
  };

  return (
    <div className="ChatPanel">
      <ChatHeader />
      <ChatBody messageList={messageList} />
      <ChatFooter onSendMessage={sendMessage} ref={textareaRef} />
    </div>
  );
};
