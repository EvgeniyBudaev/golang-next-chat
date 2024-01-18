import type { z } from "zod";
import { roomListByProfileParamsSchema } from "@/app/api/room/listByProfile/schemas";

export type TRoomListByProfileParams = z.infer<
  typeof roomListByProfileParamsSchema
>;
