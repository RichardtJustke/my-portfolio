"use client";
import { Suspense } from "react";
import Loader from "@/components/Loader";

export default function RootTemplate({ children }) {
  // fallback aparece rapidamente durante transições (streaming)
  return <Suspense fallback={<Loader fullscreen={false} />}>{children}</Suspense>;
}
