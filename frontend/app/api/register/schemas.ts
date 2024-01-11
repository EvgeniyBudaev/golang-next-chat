import { z } from "zod";

export const signupParamsSchema = z.any();

export const signupSchema = z.object({
  email: z.string(),
  firstName: z.string(),
  lastName: z.string(),
  mobileNumber: z.string(),
  password: z.string(),
  userName: z.string(),
});
