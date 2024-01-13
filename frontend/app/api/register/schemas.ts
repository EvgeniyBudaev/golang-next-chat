import { z } from "zod";
import { userListItemSchema } from "@/app/api/user/list/schemas";

export const registerParamsSchema = z.object({
  email: z.string(),
  firstName: z.string(),
  lastName: z.string(),
  mobileNumber: z.string(),
  password: z.string(),
  userName: z.string(),
});

export const registerSchema = z.object({
  data: userListItemSchema,
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
