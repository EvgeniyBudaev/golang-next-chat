"use client";

import { signOut } from "next-auth/react";
import { type FC, useEffect } from "react";
import { useSessionNext } from "@/app/shared/hooks";
import { Avatar } from "@/app/uikit/components/avatar";
import "./UserPanel.scss";
import { Icon } from "@/app/uikit/components/icon";

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

  useEffect(() => {
    if (status != "loading" && session && session?.error === "RefreshAccessTokenError") {
      signOut({ callbackUrl: "/" });
    }
  }, [session, status]);

  const handleLogout = () => {
    keycloakSessionLogOut().then(() => signOut({ callbackUrl: "/" }));
  };

  return (
    <div className="UserPanel">
      <div className="UserPanel-Avatar">
        {isSession && <Avatar size={46} user={session?.user?.name} />}
      </div>
      <div className="UserPanel-List">
        <div className="UserPanel-ListItem" onClick={handleLogout}>
          <Icon type="Exit" />
        </div>
      </div>
    </div>
  );
};
