import { useTranslation } from "@/app/i18n";
import { LoginPage } from "@/app/pages/loginPage";

export default async function LoginRoute({
  params: { lng },
}: {
  params: { lng: string };
}) {
  const { i18n } = await useTranslation(lng, "index");

  return <LoginPage i18n={i18n} />;
}
