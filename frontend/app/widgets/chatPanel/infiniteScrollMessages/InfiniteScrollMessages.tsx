"use client";

import { type FC, useEffect, useState } from "react";
import { useFormState } from "react-dom";
import { useInView } from "react-intersection-observer";
import { getMessageListByRoomAction } from "@/app/actions/message/listByRoom/getMessageListByRoomAction";
import { Message } from "@/app/entities/message";
import { DEFAULT_PAGE_LIMIT } from "@/app/shared/constants/pagination";
import { type TMessage } from "@/app/shared/types/message";
import { getFormattedMessages } from "@/app/widgets/chatPanel/utils";
import "./InfiniteScrollMessages.scss";

type TProps = {
  initialMessages: TMessage[];
  roomId?: number;
  userId?: string;
};
export const InfiniteScrollMessages: FC<TProps> = ({
  initialMessages,
  roomId,
  userId,
}) => {
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };
  const [messageList, setMessageList] = useState<TMessage[]>(initialMessages);
  const [page, setPage] = useState(1);
  const [isLoading, setIsLoading] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const [ref, inView] = useInView();
  const [state, formAction] = useFormState(
    getMessageListByRoomAction,
    initialState,
  );

  const loadMore = () => {
    if (isLoading) return;
    if (roomId) {
      setIsLoading(true);
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
    if (state?.success && userId) {
      if (state?.data?.content.length === 0) {
        setHasMore(false);
        setIsLoading(false);
        return;
      }
      const nextPage = page + 1;
      setPage(nextPage);
      setMessageList((prev: TMessage[]) => {
        const content = state?.data?.content ?? [];
        const formattedMessages = getFormattedMessages(content, userId);
        return [...(prev?.length ? prev : []), ...formattedMessages];
      });
      setIsLoading(false);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [state.data]);

  useEffect(() => {
    console.log("messageList: ", messageList);
  }, [messageList]);

  return (
    <>
      {(messageList ?? []).map((message: TMessage) => (
        <Message key={message.id} message={message} />
      ))}
      <div ref={ref}>
        <div>
          {isLoading && <p>Loading...</p>}
          {/*{!isLoading && !hasMore && <p>No more items to load.</p>}*/}
        </div>
      </div>
    </>
  );
};
