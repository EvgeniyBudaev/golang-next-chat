import type { z } from "zod";
import { paginationSchema } from "@/app/api/pagination/schemas";

export type TPagination = z.infer<typeof paginationSchema>;
