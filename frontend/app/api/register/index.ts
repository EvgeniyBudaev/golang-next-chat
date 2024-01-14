import { register } from "./domain";
import { registerResponseSchema, registerParamsSchema } from "./schemas";
import { TRegister, TRegisterParams, TRegisterResponse } from "./types";

export {
  register,
  registerParamsSchema,
  registerResponseSchema,
  type TRegister,
  type TRegisterResponse,
  type TRegisterParams,
};
