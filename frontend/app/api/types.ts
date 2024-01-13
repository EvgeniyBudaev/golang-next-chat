export type TApiFunction<TParams, TResponse> = (
  params: TParams,
) => Promise<TResponse>;

export type TSearchParams = { [key: string]: string | string[] | undefined };
