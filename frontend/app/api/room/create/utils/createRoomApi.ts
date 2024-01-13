import { fetchApi, TApiFunction } from "@/app/api";

import { EFormMethods } from "@/app/shared/enums";
import {
  TRoomCreateParams,
  TRoomCreateResponse,
} from "@/app/api/room/create/types";

export const createRoomApi: TApiFunction<
  TRoomCreateParams,
  TRoomCreateResponse
> = (params) => {
  const url = "/api/v1/room/create";
  console.log("url: ", url);
  return fetchApi<TRoomCreateResponse>(url, {
    method: EFormMethods.Post,
    body: params,
  });
};
