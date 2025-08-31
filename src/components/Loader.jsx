"use client";
import { motion } from "framer-motion";

export default function Loader({ fullscreen = true }) {
  return (
    <div
      className={`${
        fullscreen ? "fixed" : "absolute"
      } inset-0 z-[1000] grid place-items-center bg-black`}
    >
      <div className="flex flex-col items-center gap-6">
        {/* marca */}
        <div className="text-white text-xl font-semibold tracking-tight">
          Justke
        </div>

        {/* “bolinhas” com leve 3D (Acenternity style) */}
        <div className="flex items-center gap-3">
          {[0, 1, 2].map((i) => (
            <motion.span
              key={i}
              initial={{ y: 0, scale: 0.9, opacity: 0.8 }}
              animate={{ y: [0, -10, 0], scale: [0.9, 1.05, 0.9], opacity: [0.8, 1, 0.8] }}
              transition={{ duration: 0.9, repeat: Infinity, delay: i * 0.15, ease: "easeInOut" }}
              className="h-3 w-3 rounded-full bg-white"
            />
          ))}
        </div>

        {/* linha sutil */}
        <div className="h-px w-40 bg-white/10" />

        {/* tip opcional */}
        <p className="text-xs text-white/60">carregando…</p>
      </div>
    </div>
  );
}
