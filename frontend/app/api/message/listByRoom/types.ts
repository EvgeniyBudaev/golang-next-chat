import type { z } from "zod";
import {
  messageListByRoomParamsSchema,
  messageListByRoomResponseSchema,
  messageListItemByRoomSchema,
} from "@/app/api/message/listByRoom/schemas";

export type TMessageListByRoomParams = z.infer<
  typeof messageListByRoomParamsSchema
>;
export type TMessageListItemByRoom = z.infer<
  typeof messageListItemByRoomSchema
>;
export type TMessageListByRoomResponse = z.infer<
  typeof messageListByRoomResponseSchema
>;
