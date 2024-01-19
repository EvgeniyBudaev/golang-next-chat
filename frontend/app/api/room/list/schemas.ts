import { z } from "zod";

export const roomListItemSchema = z.object({
  id: z.number(),
  roomName: z.string(),
  title: z.string(),
});

export const roomListParamsSchema = z.any();

export const roomListResponseSchema = z.object({
  data: roomListItemSchema.array().optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
