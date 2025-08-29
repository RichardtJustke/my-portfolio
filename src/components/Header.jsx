"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useState } from "react";

const links = [
  { href: "/portfolio/identidade-visual", label: "Identidade Visual" },
  { href: "/portfolio/web-design", label: "Web Design" },
  { href: "/portfolio/social-media", label: "Social Media" },
];

export default function Header() {
  const pathname = usePathname();
  const [open, setOpen] = useState(false);
  const isActive = (href) => pathname === href || pathname?.startsWith(href + "/");

  return (
    <header className="sticky top-0 z-50 border-b border-white/10 bg-black/80 backdrop-blur-md">
      <div className="mx-auto flex h-16 w-full max-w-6xl items-center justify-between px-4 md:px-6">
        {/* LOGO (esquerda) */}
        <Link href="/" className="flex items-center gap-2">
          <svg width="24" height="24" viewBox="0 0 24 24" className="text-white">
            <path d="M6 4h6a4 4 0 1 1 0 8H9v8H6V4zm3 5h3a2 2 0 1 0 0-4H9v4z" fill="currentColor"/>
          </svg>
          <span className="font-bold tracking-tight">Justke</span>
        </Link>

        {/* LINKS (centro) */}
        <nav className="hidden md:flex items-center gap-4 lg:gap-8">
          {links.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              className={`rounded-full px-4 py-2 text-sm font-medium transition
                ${isActive(l.href) ? "bg-white/10 text-white" : "text-zinc-300 hover:bg-white/5 hover:text-white"}`}
            >
              {l.label}
            </Link>
          ))}
        </nav>

        {/* CTA CONTATO (direita) */}
        <div className="hidden md:block">
          <Link
            href="/contato"
            className="rounded-full bg-white px-5 py-2 text-sm font-semibold text-black transition hover:opacity-90"
          >
            Contato
          </Link>
        </div>

        {/* MENU MOBILE */}
        <button
          onClick={() => setOpen((s) => !s)}
          aria-label="Abrir menu"
          className="md:hidden inline-flex h-9 w-9 items-center justify-center rounded-full border border-white/15 text-white"
        >
          <div className={`relative h-4 w-5 transition-all ${open ? "rotate-45" : ""}`}>
            <span className={`absolute left-0 top-0 h-0.5 w-5 bg-white transition-all ${open ? "translate-y-2.5" : ""}`} />
            <span className={`absolute left-0 top-2 h-0.5 w-5 bg-white transition-all ${open ? "opacity-0" : ""}`} />
            <span className={`absolute left-0 bottom-0 h-0.5 w-5 bg-white transition-all ${open ? "-rotate-90 -translate-y-2.5" : ""}`} />
          </div>
        </button>
      </div>

      {/* Drawer mobile */}
      <div className={`md:hidden overflow-hidden border-t border-white/10 transition-[max-height] ${open ? "max-h-96" : "max-h-0"}`}>
        <nav className="flex flex-col gap-1 px-4 py-3">
          {links.map((l) => (
            <Link
              key={l.href}
              href={l.href}
              onClick={() => setOpen(false)}
              className={`rounded-lg px-3 py-2 text-sm font-medium transition 
                ${isActive(l.href) ? "bg-white/10 text-white" : "text-zinc-300 hover:bg-white/5 hover:text-white"}`}
            >
              {l.label}
            </Link>
          ))}
          <Link
            href="/contato"
            onClick={() => setOpen(false)}
            className="mt-2 rounded-full bg-white px-4 py-2 text-center text-sm font-semibold text-black"
          >
            Contato
          </Link>
        </nav>
      </div>
    </header>
  );
}
