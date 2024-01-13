import { z } from "zod";
import { zfd } from "zod-form-data";
import { EFormFields } from "@/app/entities/search/enums";

export const userGetListFormSchema = zfd.formData({
  [EFormFields.Search]: z.string().trim(),
});
