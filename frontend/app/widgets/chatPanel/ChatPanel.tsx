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
  console.log("ChatPanel session: ", session);
  const [messageList, setMessageList] = useState<TMessage[]>([]);
  const textareaRef = useRef<HTMLTextAreaElement | null>(null);
  const { conn } = useContext(WebsocketContext);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  console.log("ChatPanel messageList: ", messageList);
  console.log("ChatPanel conn: ", conn);

  useEffect(() => {
    // if (textarea.current) {
    //   autosize(textarea.current)
    // }

    if (conn === null) {
      // router.push("/");
      return;
    }

    conn.addEventListener("message", message => {
      console.log("ChatPanel conn.onmessage: ", message);
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

      console.log(
        "ChatPanel session?.user?.username: ",
        session?.user?.username,
      );
      console.log("ChatPanel m.username: ", m.username);
      session?.user?.username === m.username
        ? (m.type = "self")
        : (m.type = "recv");
      setMessageList([...messageList, m]);
    })

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
      <ChatFooter onSendMessage={sendMessage} ref={textareaRef} />
    </div>
  );
};
