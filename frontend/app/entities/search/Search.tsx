"use client";

import clsx from "clsx";
import { debounce } from "lodash";
import {
  useCallback,
  type ChangeEvent,
  type FC,
  type KeyboardEvent,
  useState,
  useRef,
  useEffect,
} from "react";
import { useFormState } from "react-dom";
import { userGetListAction } from "@/app/actions/user/list/userGetListAction";
import { TUserListItem } from "@/app/api/user/list/types";
import { EFormFields } from "@/app/entities/search/enums";
import { EFormMethods } from "@/app/shared/enums";
import { Icon } from "@/app/uikit/components/icon";
import "./Search.scss";

type TProps = {
  className?: string;
  onChangeInputValue?: (event: ChangeEvent<HTMLInputElement>) => void;
  onChangeSearchState?: (userList: TUserListItem[]) => void;
};

export const Search: FC<TProps> = ({
  className,
  onChangeInputValue,
  onChangeSearchState,
}) => {
  const DEBOUNCE_TIMEOUT = 2000;
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };

  const buttonRef = useRef<HTMLInputElement>(null);
  const [isActive, setIsActive] = useState(false);
  const [state, formAction] = useFormState(userGetListAction, initialState);

  useEffect(() => {
    if (!state) return;
    onChangeSearchState?.(state?.data as TUserListItem[]);
  }, [onChangeSearchState, state]);

  // eslint-disable-next-line react-hooks/exhaustive-deps
  const debouncedFetcher = useCallback(
    debounce((event) => {
      buttonRef.current && buttonRef.current.click();
    }, DEBOUNCE_TIMEOUT),
    [],
  );

  const handleBlur = () => {
    setIsActive(false);
  };

  const handleFocus = () => {
    setIsActive(true);
  };

  const handleKeyDown = (event: KeyboardEvent) => {
    if (event.key === "Escape") {
      setIsActive(false);
    }
  };

  const handleSubmit = (event: ChangeEvent<HTMLInputElement>) => {
    debouncedFetcher(event);
    onChangeInputValue?.(event);
  };

  return (
    <div
      className={clsx("Search", className, {
        Search__active: isActive,
      })}
    >
      <form
        action={formAction}
        className="Search-Form"
        method={EFormMethods.Post}
      >
        <Icon className="Search-Icon" type="Search" />
        <div className="Search-InputWrapper">
          <input
            autoComplete="off"
            className="Search-Input"
            name={EFormFields.Search}
            placeholder={"Search"}
            type="text"
            onBlur={handleBlur}
            onChange={handleSubmit}
            onFocus={handleFocus}
            onKeyDown={handleKeyDown}
          />
        </div>
        <input hidden={true} ref={buttonRef} type="submit" />
      </form>
    </div>
  );
};
