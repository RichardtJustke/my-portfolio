"use client";
import Image from "next/image";
import Link from "next/link";
import { useRef } from "react";

/** 🔤 Título da seção */
const sectionTitle = "Domine o Digital com Design & Código Inteligente"

function Card3D({ title, copy, img, alt, href }) {
  const ref = useRef(null);

  const onMove = (e) => {
    const el = ref.current;
    if (!el) return;
    const rect = el.getBoundingClientRect();
    const x = e.clientX - rect.left;
    const y = e.clientY - rect.top;
    const midX = rect.width / 2;
    const midY = rect.height / 2;
    const rotY = ((x - midX) / midX) * 8;
    const rotX = -((y - midY) / midY) * 8;

    el.style.transform = `rotateX(${rotX}deg) rotateY(${rotY}deg)`;
    el.style.setProperty("--px", `${x}px`);
    el.style.setProperty("--py", `${y}px`);
  };

  const onLeave = () => {
    const el = ref.current;
    if (!el) return;
    el.style.transform = `rotateX(0) rotateY(0)`;
  };

  return (
    <div className="relative [perspective:1200px]" onMouseMove={onMove} onMouseLeave={onLeave}>
      <div
        ref={ref}
        className="relative h-[360px] rounded-2xl border border-white/10 bg-white/[0.03] p-5 transition-transform duration-150 ease-linear will-change-transform"
      >
        {/* brilho */}
        <div
          className="pointer-events-none absolute inset-0 rounded-2xl opacity-0 transition-opacity duration-200 group-hover:opacity-100"
          style={{
            background:
              "radial-gradient(600px circle at var(--px) var(--py), rgba(255,255,255,.08), transparent 40%)",
          }}
        />
        {/* imagem */}
        <div className="relative h-[55%] w-full overflow-hidden rounded-xl mb-4">
          <Image src={img} alt={alt} fill className="object-cover" sizes="(min-width: 1024px) 40vw, 90vw" />
        </div>

        {/* textos */}
        <h3 className="text-lg font-bold">{title}</h3>
        <p className="mt-2 text-sm text-zinc-300 line-clamp-3">{copy}</p>

        {/* CTA */}
        <Link
          href={href}
          className="absolute bottom-5 right-5 rounded-lg bg-white px-3 py-2 text-xs font-semibold text-black hover:opacity-90 transition"
        >
          Ver projetos
        </Link>
      </div>
    </div>
  );
}

export default function WorkShowcase3D() {
  const cards = [
    {
      title: "Brand Identity",
      copy:
        "Marcas fortes, memoráveis e cheias de propósito. Uma identidade visual não é só estética, é o DNA da sua marca.",
      img: "/images/pro1.png", // brand
      alt: "Brand identity",
      href: "/portfolio/identidade-visual",   // 👈 destino
    },
    {
      title: "Web Design",
      copy:
        "Design que encanta, usabilidade que converte. Interfaces modernas, responsivas e intuitivas para sua presença digital.",
      img: "/images/pro2.png", // web design
      alt: "Web design",
      href: "/portfolio/web-design",
    },
    {
      title: "Social Media",
      copy:
        "Conteúdo que conecta, estratégia que engaja. Presença digital com propósito para gerar impacto real.",
      img: "/images/pro3.png", // social media
      alt: "Social media",
      href: "/portfolio/social-media",
    },
    {
      title: "Programação",
      copy:
        "Código inteligente que transforma ideias em sistemas. Do back-end sólido ao front-end fluido, com performance e segurança.",
      img: "/images/pro4.png", // programação
      alt: "Programação",
      href: "/portfolio/programacao",
    },
  ];

  return (
    <section className="w-full px-4 md:px-6 py-16">
      {/* título da seção centralizado e branco */}
      <div className="mx-auto max-w-7xl text-center">
        <h2 className="text-3xl md:text-5xl font-extrabold mb-12 text-white">
          {sectionTitle}
        </h2>
      </div>

      {/* grid 2x2 */}
      <div className="mx-auto max-w-7xl grid grid-cols-1 md:grid-cols-2 gap-8">
        {cards.map((c) => (
          <div key={c.title} className="group">
            <Card3D {...c} />
          </div>
        ))}
      </div>
    </section>
  );
}
