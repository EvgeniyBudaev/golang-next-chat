import { fetchApi, TApiFunction } from "@/app/api";
import { TUserListParams, TUserListResponse } from "@/app/api/user/list/types";
import { EFormMethods } from "@/app/shared/enums";

export const getUserListApi: TApiFunction<
  TUserListParams,
  TUserListResponse
> = (params) => {
  const url = `/api/v1/user/list?${new URLSearchParams(params)}`;
  console.log("url: ", url);
  return fetchApi<TUserListResponse>(url, {
    method: EFormMethods.Get,
  });
};
