"use client";

import clsx from "clsx";
import type { FC } from "react";
import { memo } from "react";

import { useDayjs } from "@/app/uikit/components/dateTime/hooks";
import {
  DATE_FORMAT,
  TIME_FORMAT,
} from "@/app/uikit/components/dateTime/constants";
import "./DateTime.scss";
import {
  ETypographyVariant,
  Typography,
} from "@/app/uikit/components/typography";

type TProps = {
  className?: string;
  classes?: { date?: string; time?: string };
  dataTestId?: string;
  isUtc?: boolean;
  isTime?: boolean;
  value?: Date | string | number | null;
};

const DateTimeComponent: FC<TProps> = ({
  className,
  classes,
  dataTestId = "uikit__date-time",
  isUtc = true,
  isTime = true,
  value,
}) => {
  const { dayjs } = useDayjs();

  return (
    <div className={clsx("DateTime", className)} data-testid={dataTestId}>
      <div className={clsx("DateTime-Date", classes?.date)}>
        <Typography
          value={
            isUtc
              ? dayjs(value).utc().format(DATE_FORMAT)
              : dayjs(value).format(DATE_FORMAT)
          }
          variant={ETypographyVariant.TextB4Regular}
        />
      </div>
      {isTime && (
        <div className={clsx("DateTime-Time", classes?.time)}>
          <Typography
            value={
              isUtc
                ? dayjs(value).utc().format(TIME_FORMAT)
                : dayjs(value).format(TIME_FORMAT)
            }
            variant={ETypographyVariant.TextB4Regular}
          />
        </div>
      )}
    </div>
  );
};

export const DateTime = memo(DateTimeComponent);
