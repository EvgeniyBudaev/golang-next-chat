import { TPagination } from "@/app/api/pagination";

type TMessageListByRoom = {
  content: TMessage[];
} & TPagination;

export type WSContent = {
  message: TMessage;
  messageListByRoom: TMessageListByRoom;
  page: number;
  limit: number;
};

export type TMessage = {
  id: string;
  roomId: number;
  userId: string;
  type: "recv" | "self" | "sys";
  createdAt: string;
  updatedAt: string;
  isDeleted: boolean;
  isEdited: boolean;
  isJoined: boolean;
  isLeft: boolean;
  profile: {
    id: string;
    firstName: string;
    lastName?: string | null;
  };
  content: string;
};
