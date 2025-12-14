"use client";
import Link from "next/link";
import { useState } from "react";
import {
  Navbar,
  NavBody,
  NavItems,
  MobileNav,
  MobileNavHeader,
  MobileNavMenu,
  MobileNavToggle,
} from "@/components/ui/resizable-navbar";

const links = [
  { name: "Identidade Visual", link: "/portfolio/identidade-visual" },
  { name: "Web Design",        link: "/portfolio/web-design" },
  { name: "Social Media",      link: "/portfolio/social-media" },
  { name: "Programação",       link: "/portfolio/programacao" },
];

export default function Header() {
  const [open, setOpen] = useState(false);

  return (
    <Navbar className="">
      {/* DESKTOP */}
      <NavBody className="hidden lg:flex">
        {/* esquerda: logo */}
        <Link href="/" className="flex items-center gap-2 text-white">
          <img src="/images/LogoW.png" alt="Justke Logo" className="h-8 w-auto" />
        </Link>

        {/* centro: itens */}
        <NavItems items={links} />

        {/* direita: CTA */}
        <Link
          href="/contato"
          className="rounded-full bg-white px-4 py-2 text-sm font-semibold text-black hover:opacity-90"
        >
          Contato
        </Link>
      </NavBody>

      {/* MOBILE */}
      <MobileNav>
        <MobileNavHeader>
          <Link href="/" className="flex items-center gap-2 text-white">
            <img src="/images/LogoW.png" alt="Justke Logo" className="h-8 w-auto" />
          </Link>
          <button onClick={() => setOpen((s) => !s)} aria-label="Abrir menu">
            <MobileNavToggle isOpen={open} onClick={() => setOpen((s) => !s)} />
          </button>
        </MobileNavHeader>

        <MobileNavMenu isOpen={open} onClose={() => setOpen(false)}>
          <nav className="flex flex-col gap-1">
            {links.map((l) => (
              <Link key={l.link} href={l.link} onClick={() => setOpen(false)}
                    className="rounded-lg px-3 py-2 text-sm text-zinc-300 hover:text-white hover:bg-white/5">
                {l.name}
              </Link>
            ))}
            <Link href="/contato" onClick={() => setOpen(false)}
                  className="mt-2 rounded-full bg-white px-4 py-2 text-center text-sm font-semibold text-black">
              Contato
            </Link>
          </nav>
        </MobileNavMenu>
      </MobileNav>
    </Navbar>
  );
}
