import { getServerSession } from "next-auth";
import { redirect } from "next/navigation";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { getRoomList } from "@/app/api/room/list/domain";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { useTranslation } from "@/app/i18n";
import { MainPage } from "@/app/pages/mainPage";
import { ERoutes } from "@/app/shared/enums";
import { createPath } from "@/app/shared/utils";

async function loader() {
  const session = await getServerSession(authOptions);
  if (!session) {
    return { isSession: false, roomList: undefined };
  }
  try {
    const [roomListResponse] = await Promise.all([getRoomList({})]);
    const roomList = roomListResponse.data as TRoomListItem[];
    return { isSession: true, roomList };
  } catch (error) {
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
  const data = await loader();
  if (!data.isSession) {
    return redirect(
      createPath({
        route: ERoutes.Login,
      }),
    );
  }
  return <MainPage roomList={data.roomList ?? []} />;
}
