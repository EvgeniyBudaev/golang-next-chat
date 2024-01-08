import { ERoutes } from "@/app/shared/enums";

type TRoutes =
  | ERoutes.Root
  | ERoutes.App
  | ERoutes.Signup
  | ERoutes.Login
  | ERoutes.PermissionDenied;

type TRoutesWithParams = ERoutes.Room;

type TCreatePathProps =
  | { route: TRoutes }
  | { route: TRoutesWithParams; params: Record<string, string | number> };

type TCreatePathPropsWithParams = Extract<TCreatePathProps, { route: any; params: any }>;

export function createPath(
  props: TCreatePathProps,
  // lng: string,
  query?: Record<string, string> | URLSearchParams,
): string {
  let path: string = props.route;

  if (props.hasOwnProperty("params")) {
    path = Object.entries((props as TCreatePathPropsWithParams).params).reduce(
      (previousValue: string, [param, value]) => previousValue.replace(`:${param}`, String(value)),
      path,
    );
  }

  if (query && Object.keys(query).length) {
    path = `${path}${path.includes("?") ? "&" : "?"}${new URLSearchParams(query)}`;
  }

  // if(lng) {
  //   path = `${lng}/${path}`;
  // }

  return path;
}
