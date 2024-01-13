import { fetchApi, TApiFunction } from "@/app/api";
import { TRegister, TRegisterParams } from "@/app/api/register/types";
import { EFormMethods } from "@/app/shared/enums";

export const registerApi: TApiFunction<TRegisterParams, TRegister> = (
  params,
) => {
  return fetchApi<TRegister>(`/api/v1/user/register`, {
    method: EFormMethods.Post,
    body: params,
  });
};
