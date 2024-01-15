import { z } from "zod";

export const messageListItemSchema = z.object({
  id: z.number(),
  roomId: z.number(),
  userId: z.string(),
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
