"use client";

import { type FC } from "react";
import { useSessionNext } from "@/app/shared/hooks";
import { Avatar } from "@/app/uikit/components/avatar";
import "./UserPanel.scss";

async function keycloakSessionLogOut() {
  try {
    await fetch(`/api/auth/logout`, { method: "GET" });
  } catch (error) {
    console.error(error);
  }
}

export const UserPanel: FC = () => {
  const { data: session, status } = useSessionNext();
  const isSession = Boolean(session);

  return (
    <div className="UserPanel">
      <div>{isSession && <Avatar size={46} user={session?.user?.name} />}</div>
    </div>
  );
};
