import { getRoomList } from "@/app/api/room/list/domain";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { useTranslation } from "@/app/i18n";
import { MainPage } from "@/app/pages/mainPage";
import { ErrorBoundary } from "@/app/shared/components/errorBoundary";
import { Layout } from "@/app/shared/components/layout";
import { getServerSession } from "next-auth";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";
import { createPath } from "@/app/shared/utils";
import { ERoutes } from "@/app/shared/enums";

async function loader() {
  // TODO: пофиксить TS у authOptions
  const session = await getServerSession(authOptions);
  const isSession = Boolean(session);
  // TODO: реализовать редирект на страницу логина, если пользователь не авторизован
  if (isSession) {
    // return redirect(
    //     createPath({
    //       route: ERoutes.Login,
    //     }),
    // );
  }
  try {
    const [roomListResponse] = await Promise.all([getRoomList({})]);
    const roomList = roomListResponse.data as TRoomListItem[];
    return { roomList };
  } catch (error) {
    console.log("Err: ", error);
    throw new Error("errorBoundary.common.unexpectedError");
  }
}

type TProps = {
  params: { lng: string };
};

export default async function MainRoute(props: TProps) {
  const { params } = props;
  const { lng } = params;
  const [{ t }] = await Promise.all([useTranslation(lng, "index")]);

  try {
    const data = await loader();
    return <Layout roomList={data.roomList} i18n={{ lng, t }} />;
  } catch (error) {
    return <ErrorBoundary i18n={{ lng, t }} message={t(error?.message)} />;
  }
}
