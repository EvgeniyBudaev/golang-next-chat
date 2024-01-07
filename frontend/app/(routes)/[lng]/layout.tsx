import type { Metadata } from "next";
import { type ReactNode } from "react";
import { WebSocketProvider } from "@/app/shared/context";

export const metadata: Metadata = {
  title: "MyChat",
  description: "My chat",
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <WebSocketProvider>
      <html lang="en">
        <body>{children}</body>
      </html>
    </WebSocketProvider>
  );
}
