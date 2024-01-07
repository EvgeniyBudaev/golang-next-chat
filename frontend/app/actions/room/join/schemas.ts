import { z } from "zod";
import { zfd } from "zod-form-data";
import { EFormFields } from "@/app/features/room/roomJoinForm/enums";
import { EMPTY_FIELD_ERROR_MESSAGE } from "@/app/shared/validation";
import { roomJoinAction } from "@/app/actions/room/join/roomJoinAction";

export const roomJoinFormSchema = zfd.formData({
  [EFormFields.RoomId]: z.string().trim().min(1, EMPTY_FIELD_ERROR_MESSAGE),
  [EFormFields.UserId]: z.string().trim().min(1, EMPTY_FIELD_ERROR_MESSAGE),
  [EFormFields.UserName]: z.string().trim().min(1, EMPTY_FIELD_ERROR_MESSAGE),
});
