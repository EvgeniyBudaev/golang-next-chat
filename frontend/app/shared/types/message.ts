export type TMessage = {
  id: number;
  roomId: string;
  userId: string;
  type: "recv" | "self";
  content: string;
};
