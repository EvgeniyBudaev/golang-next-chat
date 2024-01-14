import { z } from "zod";
import { zfd } from "zod-form-data";
import { fileSchema } from "@/app/api/upload";

const profileImageListItemSchema = z.object({
  id: z.number(),
  profileId: z.number(),
  uuid: z.string(),
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
  // image: fileSchema.or(fileSchema.array()).nullish().optional(),
});

export const profileSchema = z.object({
  uuid: z.string(),
  userId: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  images: profileImageListItemSchema.array().nullish(),
});

export const profileResponseSchema = z.object({
  data: profileSchema,
  message: z.string().optional(),
  statusCode: z.number(),
  success: z.boolean(),
});
