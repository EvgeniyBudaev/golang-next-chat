import { fetchApi, TApiFunction } from "@/app/api";
import { TRegisterResponse, TRegisterParams } from "@/app/api/register/types";
import { EFormMethods } from "@/app/shared/enums";

export const registerApi: TApiFunction<TRegisterParams, TRegisterResponse> = (
  params,
) => {
  return fetchApi<TRegisterResponse>(`/api/v1/user/register`, {
    method: EFormMethods.Post,
    body: params,
  });
};
