import { fetchApi, TApiFunction } from "@/app/api";
import { TRoomListParams, TRoomListResponse } from "@/app/api/room/list/types";
import { EFormMethods } from "@/app/shared/enums";

export const getRoomListApi: TApiFunction<
  TRoomListParams,
  TRoomListResponse
> = (params) => {
  const url = `/api/v1/ws/room/list?${new URLSearchParams(params)}`;
  console.log("url: ", url);
  return fetchApi<TRoomListResponse>(url, {
    method: EFormMethods.Get,
  });
};
