import type { z } from "zod";
import {
  roomListItemSchema,
  roomListParamsSchema,
  roomListResponseSchema,
} from "@/app/api/room/list/schemas";

export type TRoomListParams = z.infer<typeof roomListParamsSchema>;
export type TRoomListItem = z.infer<typeof roomListItemSchema>;
export type TRoomListResponse = z.infer<typeof roomListResponseSchema>;
