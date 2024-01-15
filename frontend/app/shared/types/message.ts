export type TMessage = {
  id: number;
  roomId: number;
  userId: string;
  type: "recv" | "self";
  content: string;
};
