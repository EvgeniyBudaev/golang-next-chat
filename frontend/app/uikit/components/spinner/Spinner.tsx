import type { FC } from "react";
import { Icon } from "@/app/uikit/components/icon";
import "./Spinner.scss";

type TProps = {
  dataTestId?: string;
};

export const Spinner: FC<TProps> = ({ dataTestId = "uikit__spinner" }) => {
  return (
    <div className="Spinner" data-testid={dataTestId}>
      <Icon type="Spinner" />
    </div>
  );
};
