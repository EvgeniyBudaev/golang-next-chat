import { z } from "zod";
import { roomListItemSchema } from "@/app/api/room/list/schemas";

export const roomListByProfileParamsSchema = z.object({
  profileId: z.string().or(z.number()),
});

export const roomListByProfileResponseSchema = z.object({
  data: roomListItemSchema.array().optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
