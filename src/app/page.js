import BackgroundBeams from "@/components/BackgroundBeams";

export default function Home() {
  return (
    <main className="relative min-h-[100svh] bg-black text-white overflow-hidden">
      <BackgroundBeams />
      <section className="relative z-10 flex min-h-[100svh] items-center justify-center px-6 text-center">
        <div className="max-w-3xl">
          <h1 className="text-5xl md:text-7xl font-extrabold leading-tight">
            <span className="bg-gradient-to-r from-white via-zinc-300 to-zinc-500 bg-clip-text text-transparent drop-shadow-[0_0_20px_rgba(255,255,255,.15)]">
              Justke Studio
            </span>
          </h1>

          <p className="mt-4 text-lg md:text-2xl opacity-80">
            Identidades visuais, posts criativos, design web, criatividade e código que estala na sua tela.
          </p>

          <a
            href="/portfolio"
            className="inline-block mt-8 rounded-xl px-6 py-3 bg-white text-black font-semibold
                       hover:scale-[1.02] active:scale-[0.99] transition shadow-[0_10px_30px_rgba(255,255,255,.15)]"
          >
            Ver portfólio
          </a>
        </div>
      </section>
    </main>
  );
}
