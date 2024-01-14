import type { z } from "zod";
import {
  profileCreateParamsSchema,
  profileResponseSchema,
  profileSchema,
} from "@/app/api/profile/create/schemas";

export type TProfileCreateParams = z.infer<typeof profileCreateParamsSchema>;
export type TProfile = z.infer<typeof profileSchema>;
export type TProfileResponse = z.infer<typeof profileResponseSchema>;
