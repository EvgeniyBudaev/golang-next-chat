import { z } from "zod";

const profileSchemaByMessageSchema = z.object({
  uuid: z.string(),
  firstName: z.string(),
  lastName: z.string().nullish(),
});

export const messageListItemSchema = z.object({
  uuid: z.string(),
  roomId: z.number(),
  userId: z.string(),
  type: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  isDeleted: z.boolean(),
  isEdited: z.boolean(),
  profile: profileSchemaByMessageSchema,
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
