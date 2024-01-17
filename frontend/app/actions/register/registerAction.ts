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
import { createRoom } from "@/app/api/room/create/domain";
import {
  createProfile,
  type TProfileCreateParams,
} from "@/app/api/profile/create";

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
    const userResponse = await register(mapperParams);
    if (userResponse.success) {
      const profileFormData = new FormData();
      profileFormData.append("userId", userResponse.data.id);
      profileFormData.append("username", userResponse.data.username);
      profileFormData.append("firstName", userResponse.data.firstName);
      profileFormData.append("lastName", userResponse.data.lastName);
      profileFormData.append("email", userResponse.data.email);
      profileFormData.append("isEnabled", userResponse.data.enabled.toString());
      await createProfile(profileFormData as unknown as TProfileCreateParams);
      const roomDto = {
        id: userResponse.data.id,
        userId: userResponse.data.id,
      };
      await createRoom(roomDto);
    }
    const path = createPath({
      route: ERoutes.Register,
    });
    revalidatePath(path);

    return {
      data: userResponse.data,
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
