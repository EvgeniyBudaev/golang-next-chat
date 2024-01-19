import { z } from "zod";

const profileByMessageSchema = z.object({
  id: z.string(),
  firstName: z.string(),
  lastName: z.string().nullish(),
});

export const messageListItemSchema = z.object({
  id: z.string(),
  roomId: z.number(),
  userId: z.string(),
  type: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  isDeleted: z.boolean(),
  isEdited: z.boolean(),
  profile: profileByMessageSchema,
  content: z.string(),
});

export const messageListParamsSchema = z.object({
  roomId: z.string(),
});

export const messageListResponseSchema = z.object({
  data: messageListItemSchema.array().optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
