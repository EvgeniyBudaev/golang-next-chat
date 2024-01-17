import { TSearchParams } from "@/app/api/types";

type TResponse = {
  search?: string;
  sort?: string;
};

type TMapParamsToDto = (searchParams: TSearchParams) => TResponse;

export const mapParamsToDto: TMapParamsToDto = (searchParams) => {
  const search: string | undefined = !Array.isArray(searchParams?.search)
    ? searchParams?.search ?? undefined
    : undefined;
  const sort: string | undefined = !Array.isArray(searchParams?.sort)
    ? searchParams?.sort ?? undefined
    : undefined;

  return {
    ...(search ? { search } : {}),
    ...(sort ? { sort } : {}),
  };
};
