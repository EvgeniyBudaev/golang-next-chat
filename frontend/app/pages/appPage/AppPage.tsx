"use client";

import { useRouter } from "next/navigation";
import { type FC, useContext, useEffect, useRef, useState } from "react";
import { API_URL } from "@/app/shared/constants";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { ChatBody } from "@/app/shared/components/chatBody";

export type Message = {
  content: string;
  client_id: string;
  username: string;
  room_id: string;
  type: "recv" | "self";
};

export const AppPage: FC = () => {
  const [messages, setMessage] = useState<Array<Message>>([]);
  const textarea = useRef<HTMLTextAreaElement | null>(null);
  const { conn } = useContext(WebsocketContext);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  const user = {
    id: "1",
    username: "User1",
  };

  const router = useRouter();

  // useEffect(() => {
  //   if (conn === null) {
  //     router.push("/");
  //     return;
  //   }
  //
  //   const roomId = conn.url.split("/")[5];
  //
  //   async function getUsers() {
  //     try {
  //       const res = await fetch(`${API_URL}/ws/room/${roomId}/client/list`, {
  //         method: "GET",
  //         headers: {"Content-Type": "application/json"},
  //       });
  //       const data = await res.json();
  //
  //       setUsers(data);
  //     } catch (e) {
  //       console.error(e);
  //     }
  //   }
  //
  //   getUsers();
  // }, []);

  useEffect(() => {
    // if (textarea.current) {
    //   autosize(textarea.current)
    // }

    if (conn === null) {
      router.push("/");
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
        setMessage([...messages, m]);
        return;
      }

      user?.username == m.username ? (m.type = "self") : (m.type = "recv");
      setMessage([...messages, m]);
    };

    conn.onclose = () => {};
    conn.onerror = () => {};
    conn.onopen = () => {};
  }, [textarea, messages, conn, users]);

  const sendMessage = () => {
    if (!textarea.current?.value) return;
    if (conn === null) {
      router.push("/");
      return;
    }

    if ("value" in textarea.current) {
      conn.send(textarea.current.value);
    }
    if ("value" in textarea.current) {
      textarea.current.value = "";
    }
  };

  return (
    <>
      <div>
        <div>
          <ChatBody data={messages} />
        </div>
        <div>
          <div>
            <div>
              <textarea
                ref={textarea}
                placeholder="type your message here"
                style={{ resize: "none" }}
              />
            </div>
            <div>
              <button onClick={sendMessage}>Send</button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};
