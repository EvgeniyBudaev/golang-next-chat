import { fetchApi, TApiFunction } from "@/app/api";
import {
  TMessageListByRoomParams,
  TMessageListByRoomResponse,
} from "@/app/api/message/listByRoom";
import { EFormMethods } from "@/app/shared/enums";

export const getMessageListByRoomApi: TApiFunction<
  TMessageListByRoomParams,
  TMessageListByRoomResponse
> = (params) => {
  const url = "/api/v1/room/message/list";
  console.log("url: ", url);
  return fetchApi<TMessageListByRoomResponse>(url, {
    method: EFormMethods.Post,
    body: params,
  });
};
