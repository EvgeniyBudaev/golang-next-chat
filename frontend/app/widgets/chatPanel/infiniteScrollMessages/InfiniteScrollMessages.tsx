"use client";

import { type FC, useEffect, useState } from "react";
import { useFormState } from "react-dom";
import { useInView } from "react-intersection-observer";
import { getMessageListByRoomAction } from "@/app/actions/message/listByRoom/getMessageListByRoomAction";
import { Message } from "@/app/entities/message";
import { DEFAULT_PAGE_LIMIT } from "@/app/shared/constants/pagination";
import { type TMessage } from "@/app/shared/types/message";
import "./InfiniteScrollMessages.scss";

type TProps = {
  initialMessages: TMessage[];
  roomId?: number;
};
export const InfiniteScrollMessages: FC<TProps> = ({
  initialMessages,
  roomId,
}) => {
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };
  const [messageList, setMessageList] = useState<TMessage[]>(initialMessages);
  const [page, setPage] = useState(1);
  const [ref, inView] = useInView();
  const [state, formAction] = useFormState(
    getMessageListByRoomAction,
    initialState,
  );

  const loadMore = () => {
    if (roomId) {
      const nextPage = (page + 1).toString();
      const formDataDto = new FormData();
      roomId && formDataDto.append("roomId", roomId.toString());
      formDataDto.append("page", nextPage);
      formDataDto.append("limit", DEFAULT_PAGE_LIMIT.toString());
      formAction(formDataDto);
    }
  };

  useEffect(() => {
    if (inView) {
      loadMore();
    }
  }, [inView]);

  useEffect(() => {
    setMessageList(initialMessages);
  }, [initialMessages]);

  useEffect(() => {
    if (state?.success) {
      const nextPage = page + 1;
      setPage(nextPage);
      setMessageList((prev) => [
        ...(prev?.length ? prev : []),
        ...(state?.data?.content ?? []),
      ]);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [state.data]);

  useEffect(() => {
    console.log("messageList: ", messageList);
  }, [messageList]);

  return (
    <div className="InfiniteScrollMessages">
      {(messageList ?? []).map((message: TMessage, index: number) => (
        <Message key={`${message.content}-${index}`} message={message} />
      ))}
      <div ref={ref}>
        <div>Spinner</div>
        <div>Loading...</div>
      </div>
    </div>
  );
};
