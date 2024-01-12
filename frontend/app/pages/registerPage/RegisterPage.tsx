import Link from "next/link";
import { type FC } from "react";
import { RegisterForm } from "@/app/features/register/registerForm";
import { I18nProps } from "@/app/i18n/props";
import { ERoutes } from "@/app/shared/enums";
import { createPath } from "@/app/shared/utils";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./RegisterPage.scss";

export const RegisterPage: FC<I18nProps> = ({ i18n }) => {
  return (
    <div className="RegisterPage">
      <div className="RegisterPage-Center">
        <div className="RegisterPage-Content">
          <div className="RegisterPage-Title">
            <Typography
              value={i18n.t("pages.register.title")}
              variant={ETypographyVariant.TextH1Bold}
            />
          </div>
          <RegisterForm />
          <div className="RegisterPage-HaveAccount">
            <Typography
              value={i18n.t("pages.register.haveAccount")}
              variant={ETypographyVariant.TextB3Regular}
            />
            <Link
              href={createPath({
                route: ERoutes.Login,
              })}
            >
              <Typography
                value={i18n.t("pages.register.enter")}
                variant={ETypographyVariant.TextB3Regular}
              />
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};
