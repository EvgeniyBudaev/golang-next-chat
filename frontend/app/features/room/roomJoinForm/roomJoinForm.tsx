"use client";

import { type FC, useContext } from "react";
import { useFormState } from "react-dom";
import { roomJoinAction } from "@/app/actions/room/join/roomJoinAction";
import { type TRoomListItem } from "@/app/api/room/list/types";
import { useTranslation } from "@/app/i18n/client";
import { EFormFields } from "@/app/features/room/roomJoinForm/enums";
import { WEBSOCKET_URL } from "@/app/shared/constants";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { useSessionNext } from "@/app/shared/hooks";

type TProps = {
  room: TRoomListItem;
};

export const RoomJoinForm: FC<TProps> = ({ room }) => {
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };
  const { data: session, status } = useSessionNext();
  const { t } = useTranslation("index");
  const [state, formAction] = useFormState(roomJoinAction, initialState);
  const { setConn } = useContext(WebsocketContext);

  const joinRoom = (roomId: string) => {
    const ws = new WebSocket(
      `${WEBSOCKET_URL}/room/join/${roomId}?userId=${session?.user?.id}&username=${session?.user?.username}`,
    );
    if (ws.OPEN) {
      setConn(ws);
    }
  };

  const handleSubmit = (formData: FormData) => {
    formAction(formData);
    joinRoom(room.id);
  };

  return (
    <form action={handleSubmit} className="Form">
      <input defaultValue={room.id} name={EFormFields.RoomId} type="hidden" />
      <input
        defaultValue={session?.user?.id}
        name={EFormFields.UserId}
        type="hidden"
      />
      <input
        defaultValue={session?.user?.username}
        name={EFormFields.UserName}
        type="hidden"
      />
      <button type="submit">join</button>
    </form>
  );
};
