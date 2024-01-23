import { z } from "zod";
import { zfd } from "zod-form-data";
import { EFormFields } from "@/app/widgets/chatPanel/infiniteScrollMessages/enums";

export const getMessageListByRoomFormSchema = zfd.formData({
  [EFormFields.RoomId]: z.string().trim(),
  [EFormFields.Page]: z.string().trim(),
  [EFormFields.Limit]: z.string().trim(),
});
