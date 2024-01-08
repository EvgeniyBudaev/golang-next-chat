import { type ChangeEvent, type FC, type MouseEvent } from "react";
import "./Search.scss";
import clsx from "clsx";
import { EFormMethods } from "@/app/shared/enums";
import { EFormFields } from "@/app/entities/search/enums";
import { Icon } from "@/app/uikit/components/icon";

type TProps = {
  className?: string;
  defaultSearch?: string;
  isActive?: boolean;
  onBlur?: (event: ChangeEvent<HTMLInputElement>) => void;
  onClick?: (event: MouseEvent<HTMLInputElement>) => void;
  onFocus?: (event: ChangeEvent<HTMLInputElement>) => void;
  onKeyDown?: (event: KeyboardEvent) => void;
  onSubmit?: (event: ChangeEvent<HTMLFormElement>) => void;
};

export const Search: FC<TProps> = ({
  className,
  defaultSearch,
  isActive,
  onBlur,
  onClick,
  onFocus,
  onKeyDown,
  onSubmit,
}) => {
  return (
    <div
      className={clsx("Search", className, {
        Search__active: isActive,
      })}
    >
      <form className="Search-Form" method={EFormMethods.Get} onChange={onSubmit}>
        <div className="Search-InputWrapper">
          <input
            autoComplete="off"
            className="Search-Input"
            defaultValue={defaultSearch}
            name={EFormFields.Search}
            placeholder={"Search"}
            type="text"
            onBlur={onBlur}
            onClick={onClick}
            onFocus={onFocus}
            onKeyDown={onKeyDown}
          />
        </div>
        <Icon className="Search-Icon" type="Search" />
      </form>
    </div>
  );
};
