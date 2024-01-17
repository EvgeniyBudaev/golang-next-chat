import { z } from "zod";

const profileSchema = z.object({
  uuid: z.string(),
  firstName: z.string(),
  lastName: z.string().nullish(),
});

export const roomListItemSchema = z.object({
  id: z.string(),
  uuid: z.string(),
  profile: profileSchema,
});

export const roomListParamsSchema = z.any();

export const roomListResponseSchema = z.object({
  data: roomListItemSchema.array().optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
