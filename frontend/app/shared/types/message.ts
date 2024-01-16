export type TMessage = {
  uuid: string;
  roomId: number;
  userId: string;
  type: "recv" | "self" | "sys";
  createdAt: string;
  updatedAt: string;
  isDeleted: boolean;
  isEdited: boolean;
  profile: {
    uuid: string;
    firstName: string;
    lastName?: string | null;
  };
  content: string;
};
