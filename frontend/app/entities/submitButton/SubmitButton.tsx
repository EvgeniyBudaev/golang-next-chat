"use client";

import type { FC, MouseEvent } from "react";
import { useFormStatus } from "react-dom";
import { Button } from "@/app/uikit/components/button";
import { IButtonProps } from "@/app/uikit/components/button/Button";

type TProps = {
  buttonText?: string;
  className?: string;
  onClick?: (event: MouseEvent) => void;
} & IButtonProps;

export const SubmitButton: FC<TProps> = ({
  buttonText = "",
  className,
  onClick,
  ...rest
}) => {
  const { pending } = useFormStatus();

  return (
    <Button
      aria-disabled={pending}
      className={className}
      onClick={onClick}
      type="submit"
      {...rest}
    >
      {buttonText}
    </Button>
  );
};
