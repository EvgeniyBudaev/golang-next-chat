"use server";

import { revalidatePath } from "next/cache";
import { getMessageListByRoomFormSchema } from "@/app/actions/message/listByRoom/schemas";
import {
  getMessageListByRoom,
  TMessageListByRoomResponse,
} from "@/app/api/message/listByRoom";
import { ERoutes } from "@/app/shared/enums";
import { TCommonResponseError } from "@/app/shared/types/error";
import {
  createPath,
  getErrorsResolver,
  getResponseError,
} from "@/app/shared/utils";

export async function getMessageListByRoomAction(
  prevState: any,
  formData: FormData,
) {
  console.log(
    "getMessageListByRoomAction: ",
    Object.fromEntries(formData.entries()),
  );
  const resolver = getMessageListByRoomFormSchema.safeParse(
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
    const response = (await getMessageListByRoom(
      formattedParams,
    )) as TMessageListByRoomResponse;
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
