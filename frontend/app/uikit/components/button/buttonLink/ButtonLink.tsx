import clsx from "clsx";
import Link from "next/link";
import { memo } from "react";
import type { FC, ReactNode } from "react";
import { createPath } from "@/app/shared/utils";
import "../Button.scss";

type TProps = {
  children?: ReactNode;
  className?: string;
  href: string;
};

const ButtonLinkComponent: FC<TProps> = ({ className, children, href }) => {
  return (
    <Link
      className={clsx("Button", className)}
      href={createPath({
        route: href as any,
      })}
    >
      <span>{children}</span>
    </Link>
  );
};

export const ButtonLink = memo(ButtonLinkComponent);
