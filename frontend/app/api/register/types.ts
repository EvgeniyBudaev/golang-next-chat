import type { z } from "zod";
import {
  registerParamsSchema,
  registerResponseSchema,
} from "@/app/api/register/schemas";
import { userListItemSchema } from "@/app/api/user/list/schemas";

export type TRegisterParams = z.infer<typeof registerParamsSchema>;
export type TRegister = z.infer<typeof userListItemSchema>;
export type TRegisterResponse = z.infer<typeof registerResponseSchema>;
