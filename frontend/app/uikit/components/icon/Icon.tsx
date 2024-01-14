import clsx from "clsx";
import type { FC, MouseEventHandler } from "react";
import { EColorText, EColorType } from "@/app/uikit/components/colors";
import { TColor } from "@/app/uikit/components/colors/types";
import { IconType, iconTypes } from "@/app/uikit/components/icon/iconType";
import "./Icon.scss";

const getIcon = (type: string) => iconTypes.get(type);

type TProps = {
  className?: string;
  color?: TColor;
  dataTestId?: string;
  height?: number;
  onClick?: MouseEventHandler<HTMLDivElement>;
  size?: number;
  type: IconType;
  width?: number;
};

export const Icon: FC<TProps> = ({
  className,
  color = EColorText.Dark,
  dataTestId = "uikit__icon",
  height,
  width,
  onClick,
  size,
  type,
  ...rest
}) => {
  const mainStyles = clsx(`${EColorType.Icon}-${color}`);

  return (
    <div
      className={clsx("Icon", className, mainStyles)}
      data-testid={dataTestId}
      onClick={onClick}
      {...rest}
    >
      {getIcon(type)}
    </div>
  );
};
