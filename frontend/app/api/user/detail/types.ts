import type { z } from "zod";
import {
  userDetailResponseSchema,
  userDetailSchema,
} from "@/app/api/user/detail/schemas";

export type TUserDetail = z.infer<typeof userDetailSchema>;
export type TUserDetailResponse = z.infer<typeof userDetailResponseSchema>;
