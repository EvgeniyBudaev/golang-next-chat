import { z } from "zod";
import { zfd } from "zod-form-data";
import { EMPTY_FIELD_ERROR_MESSAGE } from "@/app/shared/validation";
import { EFormFields } from "@/app/widgets/chatPanel/enums";

export const getMessageListFormSchema = zfd.formData({
  [EFormFields.RoomId]: z.string().trim().min(1, EMPTY_FIELD_ERROR_MESSAGE),
});
