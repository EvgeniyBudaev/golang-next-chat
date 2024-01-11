export type TMessage = {
  content: string;
  username: string;
  room_id: string;
  type: "recv" | "self";
};
