"use client";

import {
  type FC,
  useContext,
  useEffect,
  useMemo,
  useRef,
  useState,
} from "react";
import { useFormState } from "react-dom";
import { getMessageListAction } from "@/app/actions/message/list/getMessageListAction";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { useSessionNext } from "@/app/shared/hooks";
import { type TMessage } from "@/app/shared/types/message";
import { ChatBody } from "@/app/widgets/chatPanel/chatBody";
import { ChatFooter } from "@/app/widgets/chatPanel/chatFooter";
import { ChatHeader } from "@/app/widgets/chatPanel/chatHeader";
import { EFormFields } from "@/app/widgets/chatPanel/enums";
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
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };
  const [state, formAction] = useFormState(getMessageListAction, initialState);

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
      const m: TMessage = JSON.parse(message.data);
      if (m.content == "A new user has joined the room") {
        setUsers([...users, { userId: m.userId }]);
      }
      if (m.content == "user left the chat") {
        const deleteUser = users.filter((user) => user.userId != m.userId);
        setUsers([...deleteUser]);
        setMessageList([...messageList, m]);
        return;
      }
      session?.user?.id === m.userId ? (m.type = "self") : (m.type = "recv");
      setMessageList([...messageList, m]);
      setRoomId(Number(m.roomId));
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
    messageList,
    conn,
    users,
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
    if (state?.data) {
      return state?.data.map((message) => {
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
  }, [session?.user?.id, state?.data]);

  const handleSubmit = (formData: FormData) => {
    const formattedFormData = new FormData();
    roomId && formattedFormData.append(EFormFields.RoomId, roomId.toString());
    console.log("roomId: ", roomId);
    formAction(formattedFormData);
  };

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
      <form action={handleSubmit}>
        {/*<input*/}
        {/*  value={roomId}*/}
        {/*  hidden={true}*/}
        {/*  name={EFormFields.RoomId}*/}
        {/*  type="text"*/}
        {/*/>*/}
        <input hidden={true} ref={buttonRef} type="submit" />
      </form>
    </div>
  );
};
