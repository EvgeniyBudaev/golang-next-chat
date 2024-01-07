import type { z } from "zod";
import { roomJoinParamsSchema, roomJoinResponseSchema } from "@/app/api/room/join/schemas";

export type TRoomJoinParams = z.infer<typeof roomJoinParamsSchema>;
export type TRoomJoinResponse = z.infer<typeof roomJoinResponseSchema>;
