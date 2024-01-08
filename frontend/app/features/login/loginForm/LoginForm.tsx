"use client";

import { signIn, useSession } from "next-auth/react";
import { redirect } from "next/navigation";
import { useEffect, type FC } from "react";
import { useTranslation } from "@/app/i18n/client";
import { ERoutes } from "@/app/shared/enums";
import { createPath } from "@/app/shared/utils";
import { Button } from "@/app/uikit/components/button";
import "./LoginForm.scss";

export const LoginForm: FC = () => {
  const { data: session, status } = useSession();
  const { t } = useTranslation("index");
  const isLoading = status === "loading";

  useEffect(() => {
    if (session) {
      redirect(
        createPath({
          route: ERoutes.Root,
        }),
      );
    }
  }, [session]);

  return (
    <div className="LoginForm">
      <Button
        className="LoginForm-Button"
        onClick={() => signIn("keycloak")}
        aria-disabled={isLoading}
      >
        {t("pages.login.keycloak")}
      </Button>
    </div>
  );
};
