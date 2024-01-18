import { z } from "zod";

const accessSchema = z.object({
  impersonate: z.boolean(),
  manage: z.boolean(),
  manageGroupMembership: z.boolean(),
  mapRoles: z.boolean(),
  view: z.boolean(),
});

const attributeItemSchema = z.object({
  mobileNumber: z.string(),
});

export const userDetailSchema = z.object({
  attributes: attributeItemSchema.array(),
  createdTimestamp: z.number(),
  disableableCredentialTypes: z.string().array(),
  email: z.string(),
  emailVerified: z.boolean(),
  enabled: z.boolean(),
  firstName: z.string(),
  id: z.string(),
  lastName: z.string(),
  requiresActions: z.string().array(),
  totp: z.boolean(),
  username: z.string(),
});

export const userDetailResponseSchema = z.object({
  data: userDetailSchema.optional(),
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
