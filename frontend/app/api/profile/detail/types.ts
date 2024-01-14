import type { z } from "zod";
import { profileDetailParamsSchema } from "@/app/api/profile/detail/schemas";

export type TProfileDetailParams = z.infer<typeof profileDetailParamsSchema>;
