export const metadata = {
  title: "Portfólio – Justke",
  description: "Identidade Visual, Web Design e Social Media com estratégia e estética.",
};

export default function Portfolio() {
  return (
    <main className="min-h-screen bg-black text-white">
      {/* HERO */}
      <section className="mx-auto w-full max-w-6xl px-4 md:px-6 py-16 md:py-24">
        <h1 className="text-4xl md:text-6xl font-extrabold tracking-tight">
          Portfólio
        </h1>
        <p className="mt-4 max-w-3xl text-base md:text-lg text-zinc-300">
          Escolha uma categoria para ver meus trabalhos (copy título & subtítulo ficam aqui). 
          Projetos com foco em impacto visual, estratégia e conversão.
        </p>
      </section>

      {/* CARDS DE CATEGORIA */}
      <section className="mx-auto w-full max-w-6xl px-4 md:px-6 pb-24">
        <div className="grid gap-6 md:grid-cols-3">
          {/* Identidade Visual */}
          <a
            href="/portfolio/identidade-visual"
            className="group rounded-2xl bg-white/5 p-6 transition hover:bg-white/10"
          >
            <h2 className="text-xl font-semibold">Identidade Visual / Brand</h2>
            <p className="mt-2 text-sm text-zinc-300">
              Logos, sistemas visuais, brand books, diretrizes de uso.
            </p>
            <div className="mt-6 h-24 rounded-xl bg-gradient-to-br from-zinc-200/10 to-white/5 group-hover:from-white/15 group-hover:to-white/10 transition" />
          </a>

          {/* Web Design */}
          <a
            href="/portfolio/web-design"
            className="group rounded-2xl bg-white/5 p-6 transition hover:bg-white/10"
          >
            <h2 className="text-xl font-semibold">Web Design</h2>
            <p className="mt-2 text-sm text-zinc-300">
              Sites, landings, UI/UX e performance.
            </p>
            <div className="mt-6 h-24 rounded-xl bg-gradient-to-br from-zinc-200/10 to-white/5 group-hover:from-white/15 group-hover:to-white/10 transition" />
          </a>

          {/* Social Media */}
          <a
            href="/portfolio/social-media"
            className="group rounded-2xl bg-white/5 p-6 transition hover:bg-white/10"
          >
            <h2 className="text-xl font-semibold">Social Media</h2>
            <p className="mt-2 text-sm text-zinc-300">
              Posts, carrosséis, ads e motion.
            </p>
            <div className="mt-6 h-24 rounded-xl bg-gradient-to-br from-zinc-200/10 to-white/5 group-hover:from-white/15 group-hover:to-white/10 transition" />
          </a>
        </div>
      </section>
    </main>
  );
}
