import { z } from "zod";
import { zfd } from "zod-form-data";
import { fileSchema } from "@/app/api/upload";

const profileImageListItemSchema = z.object({
  id: z.number(),
  profileId: z.number(),
  name: z.string(),
  url: z.string(),
  size: z.number(),
  createdAt: z.string(),
  updatedAt: z.string(),
  isDeleted: z.boolean(),
  isEnabled: z.boolean(),
});

export const profileCreateParamsSchema = zfd.formData({
  userId: zfd.text(),
  username: zfd.text(),
  firstName: zfd.text(),
  lastName: zfd.text(),
  email: zfd.text(),
  isEnabled: zfd.text(),
  // image: fileSchema.or(fileSchema.array()).nullish().optional(),
});

export const profileSchema = z.object({
  id: z.string(),
  userId: z.string(),
  username: z.string(),
  firstName: z.string(),
  lastName: z.string(),
  email: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  isDeleted: z.boolean(),
  isEnabled: z.boolean(),
  images: profileImageListItemSchema.array().nullish(),
});

export const profileResponseSchema = z.object({
  data: profileSchema,
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
