import { fetchApi, TApiFunction } from "@/app/api";
import { TProfileDetailParams } from "@/app/api/profile/detail/types";
import { TProfileResponse } from "@/app/api/profile/create";
import { EFormMethods } from "@/app/shared/enums";

export const getProfileDetailApi: TApiFunction<
  TProfileDetailParams,
  TProfileResponse
> = (params) => {
  const { uuid } = params;
  const url = "/api/v1/profile/detail";
  return fetchApi<TProfileResponse>(url, {
    method: EFormMethods.Post,
    body: params,
  });
};
