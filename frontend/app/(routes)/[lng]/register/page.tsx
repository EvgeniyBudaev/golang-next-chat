import { useTranslation } from "@/app/i18n";
import { RegisterPage } from "@/app/pages/registerPage";

export default async function RegisterRoute({ params: { lng } }: { params: { lng: string } }) {
  const { t } = await useTranslation(lng, "index");

  return <RegisterPage i18n={{ lng, t }} />;
}
