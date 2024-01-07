"use client";

import { useRouter } from "next/navigation";
import { useEffect, type FC, useState, useContext } from "react";
import { useFormState, useFormStatus } from "react-dom";
import { roomJoinAction } from "@/app/actions/room/join/roomJoinAction";
import { useTranslation } from "@/app/i18n/client";
import { EFormFields } from "@/app/features/room/roomJoinForm/enums";
import { WebsocketContext } from "@/app/shared/context/webSocketContext";
import { WEBSOCKET_URL } from "@/app/shared/constants";
import { ERoutes } from "@/app/shared/enums";
import { createPath } from "@/app/shared/utils";

export const RoomJoinForm: FC = () => {
  const { t } = useTranslation("index");
  const [state, formAction] = useFormState(roomJoinAction, {});
  const user = {
    id: "1",
    username: "User1",
  };

  const { setConn } = useContext(WebsocketContext);
  const router = useRouter();

  const joinRoom = (roomId: string) => {
    const ws = new WebSocket(
      `${WEBSOCKET_URL}/ws/room/join/${roomId}?userId=${user.id}&username=${user.username}`,
    );
    if (ws.OPEN) {
      setConn(ws);
      const path = createPath({
        route: ERoutes.App,
      });
      router.push(path);
      return;
    }
  };

  const handleSubmit = (formData: FormData) => {
    formAction(formData);
    joinRoom("1");
  };

  return (
    <form action={handleSubmit} className="Form">
      <input defaultValue={"1"} name={EFormFields.RoomId} type="hidden" />
      <input defaultValue={user.id} name={EFormFields.UserId} type="hidden" />
      <input defaultValue={user.username} name={EFormFields.UserName} type="hidden" />
      <button type="submit">join</button>
    </form>
  );
};
