import type { z } from "zod";
import { roomCreateParamsSchema, roomCreateResponseSchema } from "@/app/api/room/create/schemas";

export type TRoomCreateParams = z.infer<typeof roomCreateParamsSchema>;
export type TRoomCreateResponse = z.infer<typeof roomCreateResponseSchema>;
