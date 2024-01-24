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

export enum ETypeMessage {
  Recv = "recv",
  Self = "self",
  Sys = "sys",
}

export type TMessage = {
  id: number;
  roomId: number;
  userId: string;
  type: ETypeMessage;
  createdAt: string;
  updatedAt: string;
  isDeleted: boolean;
  isEdited: boolean;
  isJoined: boolean;
  isLeft: boolean;
  profile: {
    id: number;
    firstName: string;
    lastName?: string | null;
  };
  content: string;
};
