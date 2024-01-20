import type { z } from "zod";
import {
  profileListParamsSchema,
  profileListResponseSchema,
} from "@/app/api/profile/list/schemas";
import { TProfile } from "@/app/api/profile/create";

export type TProfileListParams = z.infer<typeof profileListParamsSchema>;
export type TProfileListItem = TProfile;
export type TProfileListResponse = z.infer<typeof profileListResponseSchema>;
