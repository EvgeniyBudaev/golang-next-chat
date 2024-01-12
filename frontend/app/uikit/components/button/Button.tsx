"use client";

import clsx from "clsx";
import { memo, type FC, type DOMAttributes } from "react";
import "./Button.scss";

export interface IButtonProps extends DOMAttributes<HTMLButtonElement> {
  className?: string;
  dataTestId?: string;
  isDisabled?: boolean;
  type?: "submit" | "reset" | "button";
}

const ButtonComponent: FC<IButtonProps> = ({
  className,
  children,
  dataTestId = "uikit__button",
  isDisabled = false,
  type,
  ...rest
}) => {
  return (
    <button
      className={clsx("Button", className, {
        Button__disabled: isDisabled,
      })}
      data-testid={dataTestId}
      disabled={isDisabled}
      type={type}
      {...rest}
    >
      <span>{children}</span>
    </button>
  );
};

export const Button = memo(ButtonComponent);
