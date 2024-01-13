import { TRegisterParams } from "@/app/api/register";

type TProps = TRegisterParams & { passwordConfirm: string };

export const mapRegisterToDto = (props: TProps) => {
  const { passwordConfirm, ...form } = props;
  return form;
};
