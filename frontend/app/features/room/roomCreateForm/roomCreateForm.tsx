"use client";

import { useEffect, type FC, useState } from "react";
import { useFormState, useFormStatus } from "react-dom";
import { roomCreateAction } from "@/app/actions/room/create/roomCreateAction";
import { useTranslation } from "@/app/i18n/client";
import { EFormFields } from "@/app/features/room/roomCreateForm/enums";

export const RoomCreateForm: FC = () => {
  const { t } = useTranslation("index");
  const [state, formAction] = useFormState(roomCreateAction, {});
  const [roomName, setRoomName] = useState("");

  return (
    <form action={formAction} className="Form">
      <input defaultValue={"1"} name={EFormFields.Id} type="hidden" />
      <input
        name={EFormFields.Name}
        onChange={(e) => setRoomName(e.target.value)}
        placeholder="room name"
        type="text"
        value={roomName}
      />
      <button type="submit">create room</button>
    </form>
  );
};
