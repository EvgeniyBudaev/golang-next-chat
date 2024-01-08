import { getRoomList } from "@/app/api/room/list/domain";
import type { TRoomListItem } from "@/app/api/room/list/types";
import { useTranslation } from "@/app/i18n";
import { MainPage } from "@/app/pages/mainPage";
import { ErrorBoundary } from "@/app/shared/components/errorBoundary";
import { Layout } from "@/app/shared/components/layout";

async function loader() {
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
    return (
      <Layout i18n={{ lng, t }}>
        <MainPage roomList={data.roomList} />
      </Layout>
    );
  } catch (error) {
    return <ErrorBoundary i18n={{ lng, t }} message={t(error?.message)} />;
  }
}
