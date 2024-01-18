import { fetchApi, TApiFunction } from "@/app/api";
import { TRoomListResponse } from "@/app/api/room/list/types";
import { TRoomListByProfileParams } from "@/app/api/room/listByProfile";
import { EFormMethods } from "@/app/shared/enums";

export const getRoomListByProfileApi: TApiFunction<
  TRoomListByProfileParams,
  TRoomListResponse
> = (params) => {
  const url = "/api/v1/profile/room/list";
  console.log("getRoomListByProfileApi url: ", url);
  return fetchApi<TRoomListResponse>(url, {
    method: EFormMethods.Post,
    body: params,
  });
};
