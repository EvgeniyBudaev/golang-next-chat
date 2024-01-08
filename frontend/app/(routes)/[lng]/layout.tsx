import { dir } from "i18next";
import type { Metadata } from "next";
import { type ReactNode } from "react";
import { useTranslation } from "@/app/i18n";
import { I18nContextProvider } from "@/app/i18n/context";
import { Layout } from "@/app/shared/components/layout";
import { WebSocketProvider } from "@/app/shared/context";
import { SessionProviderWrapper } from "@/app/shared/utils/auth";

export const metadata: Metadata = {
  title: "MyChat",
  description: "My chat",
};

export default async function RootLayout({
  children,
  params: { lng },
}: {
  children: ReactNode;
  params: { lng: string };
}) {
  const { t } = await useTranslation(lng, "index");

  return (
    <WebSocketProvider>
      <SessionProviderWrapper>
        <I18nContextProvider lng={lng}>
          <html lang={lng} dir={dir(lng)}>
            <body>
              {/*<Layout i18n={{ lng, t }}>{children}</Layout>*/}
              {children}
            </body>
          </html>
        </I18nContextProvider>
      </SessionProviderWrapper>
    </WebSocketProvider>
  );
}
