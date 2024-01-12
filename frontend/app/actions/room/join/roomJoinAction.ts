"use server";

import { revalidatePath } from "next/cache";
import { joinRoom } from "@/app/api/room/join/domain";
import { roomJoinFormSchema } from "@/app/actions/room/join/schemas";
import { ERoutes } from "@/app/shared/enums";
import {
  createPath,
  getErrorsResolver,
  getResponseError,
} from "@/app/shared/utils";
import { TCommonResponseError } from "@/app/shared/types/error";

export async function roomJoinAction(prevState: any, formData: FormData) {
  const resolver = roomJoinFormSchema.safeParse(
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
    // const response = await joinRoom(formattedParams);
    // console.log("response: ", response);
    // const path = createPath({
    //   route: ERoutes.App,
    // });
    // revalidatePath(path);
    // return {data: response.data, error: undefined, errors: undefined, success: true};

    return {
      data: undefined,
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
