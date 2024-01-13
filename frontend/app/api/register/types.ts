import type { z } from "zod";
import {
  registerParamsSchema,
  registerSchema,
} from "@/app/api/register/schemas";

export type TRegisterParams = z.infer<typeof registerParamsSchema>;
export type TRegister = z.infer<typeof registerSchema>;
