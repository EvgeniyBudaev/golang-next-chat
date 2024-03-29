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
  useMemo,
} from "react";
import { useFormState } from "react-dom";
import { getProfileListAction } from "@/app/actions/profile/list/getProfileListAction";
import { type TProfileListItem } from "@/app/api/profile/list";
import { EFormFields } from "@/app/entities/search/enums";
import { useSessionNext } from "@/app/shared/hooks";
import { Icon } from "@/app/uikit/components/icon";
import "./Search.scss";

type TProps = {
  className?: string;
  onChangeInputValue?: (event: ChangeEvent<HTMLInputElement>) => void;
  onChangeSearchState?: (list: TProfileListItem[]) => void;
};

export const Search: FC<TProps> = ({
  className,
  onChangeInputValue,
  onChangeSearchState,
}) => {
  const DEBOUNCE_TIMEOUT = 100;
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };
  const { data: session, status } = useSessionNext();
  const buttonRef = useRef<HTMLInputElement>(null);
  const [isActive, setIsActive] = useState(false);
  const [state, formAction] = useFormState(getProfileListAction, initialState);

  const listWithoutSessionUser = useMemo(() => {
    const list = state?.data as TProfileListItem[];
    return (list ?? []).filter((item) => {
      return item.username !== session?.user?.username;
    });
  }, [state?.data, session?.user?.username]);

  useEffect(() => {
    if (!state) return;
    onChangeSearchState?.(listWithoutSessionUser);
  }, [listWithoutSessionUser, state]);

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
      <form action={formAction} className="Search-Form">
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
