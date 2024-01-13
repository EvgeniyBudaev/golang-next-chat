import { TSearchParams } from "@/app/api/types";

export const mapParamsToDto = (searchParams: TSearchParams) => {
  const search = !Array.isArray(searchParams?.search)
    ? searchParams?.search ?? undefined
    : undefined;
  const sort = !Array.isArray(searchParams?.sort)
    ? searchParams?.sort ?? undefined
    : undefined;

  return {
    ...(search ? { search: searchParams?.search } : {}),
    ...(sort ? { sort: searchParams?.sort } : {}),
  };
};
