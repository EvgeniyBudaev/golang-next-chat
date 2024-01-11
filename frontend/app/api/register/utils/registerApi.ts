import { fetchApi, TApiFunction } from "@/app/api";
import { TSignup, TSignupParams } from "@/app/api/register/types";
import { EFormMethods } from "@/app/shared/enums";

export const registerApi: TApiFunction<TSignupParams, TSignup> = (params) => {
  return fetchApi<TSignup>(`/api/v1/user/register`, {
    method: EFormMethods.Post,
    body: params,
  });
};
