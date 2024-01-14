import { fetchApi, TApiFunction } from "@/app/api";
import {
  TProfileCreateParams,
  TProfileResponse,
} from "@/app/api/profile/create/types";
import { EFormMethods } from "@/app/shared/enums";

export const createProfileApi: TApiFunction<
  TProfileCreateParams,
  TProfileResponse
> = (params) => {
  console.log("[createProfileApi url] ", "/api/v1/profile/create");
  return fetchApi<TProfileResponse>(`/api/v1/profile/create`, {
    method: EFormMethods.Post,
    body: params,
  });
};
