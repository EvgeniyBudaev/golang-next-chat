import type { z } from "zod";
import {
  messageListItemSchema,
  messageListParamsSchema,
  messageListResponseSchema,
} from "@/app/api/message/list/schemas";

export type TMessageListParams = z.infer<typeof messageListParamsSchema>;
export type TMessageListItem = z.infer<typeof messageListItemSchema>;
export type TMessageListResponse = z.infer<typeof messageListResponseSchema>;
