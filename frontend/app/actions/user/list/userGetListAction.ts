"use server";

import { revalidatePath } from "next/cache";
import { ERoutes } from "@/app/shared/enums";
import { TCommonResponseError } from "@/app/shared/types/error";
import {
  createPath,
  getErrorsResolver,
  getResponseError,
} from "@/app/shared/utils";
import { userGetListFormSchema } from "@/app/actions/user/list/schemas";
import { getUserList } from "@/app/api/user/list/domain";
import { mapParamsToDto } from "@/app/api/user/list/utils";

export async function userGetListAction(prevState: any, formData: FormData) {
  console.log("[userGetListAction] ", userGetListAction);
  const resolver = userGetListFormSchema.safeParse(
    Object.fromEntries(formData.entries()),
  );

  if (!resolver.success) {
    const errors = getErrorsResolver(resolver);
    return {
      data: undefined,
      error: undefined,
      errors: errors,
      success: false,
    };
  }

  try {
    const formattedParams = {
      ...resolver.data,
    };
    const paramsToDto = mapParamsToDto(formattedParams);
    const response = await getUserList(paramsToDto);
    const path = createPath({
      route: ERoutes.Root,
    });
    revalidatePath(path);

    return {
      data: response.data,
      error: undefined,
      errors: undefined,
      success: true,
    };
  } catch (error) {
    const errorResponse = error as Response;
    const responseData: TCommonResponseError = await errorResponse.json();
    const { message: formError, fieldErrors } =
      getResponseError(responseData) ?? {};

    return {
      data: undefined,
      error: formError,
      errors: fieldErrors,
      success: false,
    };
  }
}
