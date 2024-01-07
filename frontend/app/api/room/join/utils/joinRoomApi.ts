import { fetchApi, TApiFunction } from "@/app/api";

import { EFormMethods } from "@/app/shared/enums";
import { TRoomJoinParams, TRoomJoinResponse } from "@/app/api/room/join/types";

export const joinRoomApi: TApiFunction<TRoomJoinParams, TRoomJoinResponse> = (params) => {
  const url = `/api/v1/ws/room/join/${params.roomId}?userId=${params.userId}&username=${params.username}`;
  console.log("url: ", url);
  return fetchApi<TRoomJoinResponse>(url, {
    method: EFormMethods.Post,
    body: params,
  });
};
