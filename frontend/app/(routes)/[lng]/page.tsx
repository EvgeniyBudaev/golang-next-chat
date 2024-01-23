import isEmpty from "lodash/isEmpty";
import { getServerSession } from "next-auth";
import { redirect } from "next/navigation";
import { authOptions } from "@/app/api/auth/[...nextauth]/route";
import { getProfileDetail } from "@/app/api/profile/detail";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { getRoomListByProfile } from "@/app/api/room/listByProfile";
import { useTranslation } from "@/app/i18n";
import { MainPage } from "@/app/pages/mainPage";
import { ERoutes } from "@/app/shared/enums";
import { TSession } from "@/app/shared/types/session";
import { createPath } from "@/app/shared/utils";

async function loader() {
  const session = (await getServerSession(authOptions)) as TSession;
  if (!session) {
    return redirect(
      createPath({
        route: ERoutes.Login,
      }),
    );
  }
  try {
    const profileResponse = await getProfileDetail({
      username: session?.user?.username,
    });
    const roomListResponse = await getRoomListByProfile({
      profileId: profileResponse.data.id.toString(),
    });
    const roomListAllByProfile = roomListResponse.data as TRoomListItem[];
    const roomListByProfile = !isEmpty(roomListAllByProfile)
      ? roomListAllByProfile.filter(
          (room) => room?.roomName !== session?.user?.username,
        )
      : [];
    return { isSession: true, roomListByProfile };
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
  return <MainPage roomListByProfile={data.roomListByProfile ?? []} />;
}
