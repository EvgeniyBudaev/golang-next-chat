import { fetchApi, TApiFunction } from "@/app/api";
import {
  TMessageListParams,
  TMessageListResponse,
} from "@/app/api/message/list/types";
import { EFormMethods } from "@/app/shared/enums";

export const getMessageListApi: TApiFunction<
  TMessageListParams,
  TMessageListResponse
> = (params) => {
  const url = "/api/v1/room/message/detail";
  console.log("getMessageListApi url: ", url);
  return fetchApi<TMessageListResponse>(url, {
    method: EFormMethods.Post,
    body: params,
  });
};
