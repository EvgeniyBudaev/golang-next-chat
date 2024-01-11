import type { z } from "zod";
import { signupParamsSchema, signupSchema } from "@/app/api/register/schemas";

export type TSignupParams = z.infer<typeof signupParamsSchema>;
export type TSignup = z.infer<typeof signupSchema>;
