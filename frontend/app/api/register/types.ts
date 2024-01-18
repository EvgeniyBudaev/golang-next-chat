import type { z } from "zod";
import {
  registerParamsSchema,
  registerResponseSchema,
} from "@/app/api/register/schemas";
import { userDetailSchema } from "@/app/api/user/detail";

export type TRegisterParams = z.infer<typeof registerParamsSchema>;
export type TRegister = z.infer<typeof userDetailSchema>;
export type TRegisterResponse = z.infer<typeof registerResponseSchema>;
