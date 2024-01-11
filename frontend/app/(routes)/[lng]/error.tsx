"use client";

import { useTranslation } from "@/app/i18n/client";
import { ErrorBoundary } from "@/app/shared/components/errorBoundary";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  const { i18n, t } = useTranslation("index");
  const errorMessage = t("errorBoundary.common.unexpectedError");

  return <ErrorBoundary i18n={i18n} message={errorMessage} />;
}
