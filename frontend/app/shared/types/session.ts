import { EPermissions } from "@/app/shared/enums";

export type TSession = {
  access_token: string;
  error?: string | null;
  expires: string;
  id_token: string;
  roles: EPermissions[];
  user: {
    name: string;
    email: string;
    id: string;
    username: string;
  };
};
