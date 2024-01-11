import { type FC } from "react";
import { LoginForm } from "@/app/features/login/loginForm";
import { I18nProps } from "@/app/i18n/props";
import { ETypographyVariant, Typography } from "@/app/uikit/components/typography";
import "./LoginPage.scss";
import { ERoutes } from "@/app/shared/enums";
import { createPath } from "@/app/shared/utils";
import Link from "next/link";

export const LoginPage: FC<I18nProps> = ({ i18n }) => {
  return (
    <div className="LoginPage">
      <div className="LoginPage-Center">
        <div className="LoginPage-Content">
          <div className="LoginPage-Title">
            <Typography
              value={i18n.t("pages.login.title")}
              variant={ETypographyVariant.TextH3Bold}
            />
          </div>
          <div className="LoginPage-Title">
            <Typography
              value={i18n.t("pages.login.description")}
              variant={ETypographyVariant.TextB3Regular}
            />
          </div>
          <LoginForm />
          <div className="LoginPage-NoAccount">
            <Typography
              value={i18n.t("pages.login.noAccount")}
              variant={ETypographyVariant.TextB3Regular}
            />
            <Link
              href={createPath({
                route: ERoutes.Register,
              })}
            >
              <Typography
                value={i18n.t("pages.login.goToSignup")}
                variant={ETypographyVariant.TextB3Regular}
              />
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};
