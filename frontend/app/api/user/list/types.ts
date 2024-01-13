import type { z } from "zod";
import {
  userListItemSchema,
  userListParamsSchema,
  userListResponseSchema,
} from "@/app/api/user/list/schemas";

export type TUserListParams = z.infer<typeof userListParamsSchema>;
export type TUser = z.infer<typeof userListItemSchema>;
export type TUserListResponse = z.infer<typeof userListResponseSchema>;
