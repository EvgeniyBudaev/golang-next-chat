import { z } from "zod";
import { paginationSchema } from "@/app/api/pagination";

const profileSchema = z.object({
  id: z.number(),
  firstName: z.string(),
  lastName: z.string(),
});

export const messageListItemByRoomSchema = z.object({
  id: z.number(),
  roomId: z.number(),
  userId: z.string(),
  type: z.enum(["recv", "self", "sys"]),
  createdAt: z.string(),
  updatedAt: z.string(),
  isDeleted: z.boolean(),
  isEdited: z.boolean(),
  isJoined: z.boolean(),
  isLeft: z.boolean(),
  profile: profileSchema,
  content: z.string(),
});

export const messageListByRoomParamsSchema = z.any();

export const messageListByRoomSchema = paginationSchema.extend({
  content: messageListItemByRoomSchema.array(),
});

export const messageListByRoomResponseSchema = z.object({
  data: messageListByRoomSchema.optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
