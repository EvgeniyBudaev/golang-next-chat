import { z } from "zod";
import { profileSchema } from "@/app/api/profile/create";

export const profileListParamsSchema = z.any();

export const profileListResponseSchema = z.object({
  data: profileSchema.array().optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
