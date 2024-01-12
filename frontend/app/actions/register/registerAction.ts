"use server";

import { revalidatePath } from "next/cache";
import { signupFormSchema } from "@/app/actions/register/schemas";
import { register } from "@/app/api/register/domain";
import { mapRegisterToDto } from "@/app/api/register/utils";
import { TCommonResponseError } from "@/app/shared/types/error";
import {
  getResponseError,
  getErrorsResolver,
  normalizePhoneNumber,
  createPath,
} from "@/app/shared/utils";
import { ERoutes } from "@/app/shared/enums";

export async function registerAction(prevState: any, formData: FormData) {
  const resolver = signupFormSchema.safeParse(
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
      mobileNumber: normalizePhoneNumber(resolver.data?.mobileNumber),
    };
    const mapperParams = mapRegisterToDto(formattedParams);
    console.log("[mapperParams] ", mapperParams);
    const response = await register(mapperParams);
    console.log("[response] ", response);
    const path = createPath({
      route: ERoutes.Register,
    });
    revalidatePath(path);
    return {
      data: response,
      error: undefined,
      errors: undefined,
      success: true,
    };
  } catch (error) {
    const errorResponse = error as Response;
    const responseData: TCommonResponseError = await errorResponse.json();
    const { message: formError, fieldErrors } =
      getResponseError(responseData) ?? {};
    console.log("[formError] ", formError);
    console.log("[fieldErrors] ", fieldErrors);
    return {
      data: undefined,
      error: formError,
      errors: fieldErrors,
      success: false,
    };
  }
}
