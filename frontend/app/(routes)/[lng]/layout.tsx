import { dir } from "i18next";
import type { Metadata } from "next";
import { type ReactNode } from "react";
import { useTranslation } from "@/app/i18n";
import { I18nContextProvider } from "@/app/i18n/context";
import { WebSocketProvider } from "@/app/shared/context";
import { SessionProviderWrapper } from "@/app/shared/utils/auth";
import { ToastContainer } from "@/app/uikit/components/toast/toastContainer";
import "react-toastify/dist/ReactToastify.css";

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
              {children}
              <ToastContainer />
            </body>
          </html>
        </I18nContextProvider>
      </SessionProviderWrapper>
    </WebSocketProvider>
  );
}
