import { z } from "zod";
import { zfd } from "zod-form-data";
import { fileSchema } from "@/app/api/upload";
import { profileSchema } from "@/app/api/profile/create";

export const profileDetailParamsSchema = z.object({
  username: z.string(),
});
