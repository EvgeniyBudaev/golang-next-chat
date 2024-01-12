"use client";

import { memo, type FC } from "react";
import { useTranslation } from "@/app/i18n/client";
import { FadeIn } from "@/app/uikit/components/fadeIn";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";
import "./Error.scss";

type TProps = {
  errors?: string[] | string;
};

const ErrorComponent: FC<TProps> = ({ errors }) => {
  const { t } = useTranslation("index");

  return (
    <ul className="Error-List">
      {errors &&
        Array.isArray(errors) &&
        errors.map((error, index) => (
          <li className="Error-ListItem" key={`error-item-${index}`}>
            <FadeIn>
              <Typography
                value={t(error)}
                variant={ETypographyVariant.TextB3Regular}
              />
            </FadeIn>
          </li>
        ))}
      {errors && !Array.isArray(errors) && (
        <li className="Error-ListItem">
          <FadeIn>
            <Typography
              value={t(errors)}
              variant={ETypographyVariant.TextB3Regular}
            />
          </FadeIn>
        </li>
      )}
    </ul>
  );
};

export const Error = memo(ErrorComponent);
