import AboutSection from "@/components/AboutSection";
import InfiniteLogos from "@/components/InfiniteLogos"; // <-- importa a seção das logos

export const metadata = {
  title: "Portfólio – Justke",
  description: "Design bonito e código inteligente para sua marca.",
};

export default function Portfolio() {
  return (
    <main className="min-h-screen bg-black text-white flex flex-col items-center px-6">
      {/* HERO central */}
      <section className="text-center max-w-4xl pt-24">
        <h1 className="text-4xl md:text-6xl font-extrabold leading-tight">
          Design que encanta. Código que gera resultado.
        </h1>
        <p className="mt-6 text-lg md:text-xl text-zinc-300">
          Transformo ideias em marcas fortes e experiências digitais inteligentes que conquistam clientes.
        </p>
      </section>

      {/* SERVIÇOS simples */}
      <section className="mt-16 grid gap-6 md:grid-cols-4 text-center max-w-6xl w-full">
        <a
          href="/portfolio/identidade-visual"
          className="rounded-xl bg-white/5 px-6 py-10 hover:bg-white/10 transition"
        >
          <h2 className="text-lg font-semibold">Identidade Visual</h2>
        </a>
        <a
          href="/portfolio/web-design"
          className="rounded-xl bg-white/5 px-6 py-10 hover:bg-white/10 transition"
        >
          <h2 className="text-lg font-semibold">Web Design</h2>
        </a>
        <a
          href="/portfolio/social-media"
          className="rounded-xl bg-white/5 px-6 py-10 hover:bg-white/10 transition"
        >
          <h2 className="text-lg font-semibold">Social Media</h2>
        </a>
        <a
          href="/portfolio/programacao"
          className="rounded-xl bg-white/5 px-6 py-10 hover:bg-white/10 transition"
        >
          <h2 className="text-lg font-semibold">Programação</h2>
        </a>
      </section>

      {/* SOBRE MIM */}
      <AboutSection />

      {/* LOGOS JA FEITAS */}
      <InfiniteLogos />

      <div className="h-20" />
    </main>
  );
}
