import { z } from "zod";
import { userDetailSchema } from "@/app/api/user/detail";

export const registerParamsSchema = z.object({
  email: z.string(),
  firstName: z.string(),
  lastName: z.string(),
  mobileNumber: z.string(),
  password: z.string(),
  userName: z.string(),
});

export const registerResponseSchema = z.object({
  data: userDetailSchema,
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
