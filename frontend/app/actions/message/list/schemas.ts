import { z } from "zod";
import { EFormFields } from "@/app/widgets/chatPanel/enums";

export const messageGetListFormSchema = z.object({
  [EFormFields.RoomId]: z.string(),
});
