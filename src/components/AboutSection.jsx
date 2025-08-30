import Image from "next/image";

export default function AboutSection() {
  return (
    <section className="mx-auto w-full max-w-6xl px-4 md:px-6 py-20">
      <div className="grid items-center gap-10 md:grid-cols-2">
        {/* Coluna ESQUERDA: título, subtítulo, texto */}
        <div>
          <h2 className="text-3xl md:text-5xl font-extrabold leading-tight">
            Designer & Programador — Criando Marcas e Códigos que Conectam
          </h2>


          <p className="mt-6 text-zinc-300 leading-relaxed">
            Com 4 anos de experiência em design e estudando Ciência da Computação,
            atuo desenvolvendo identidades visuais estratégicas, sites modernos e
            soluções digitais completas. Minha missão é dar vida a marcas com
            estética profissional e código inteligente — criando projetos que não
            só encantam visualmente, mas também funcionam de forma impecável. Seja
            para fortalecer sua presença online, lançar um novo negócio ou
            modernizar sua marca, eu ajudo você a se destacar no mercado com
            impacto e autenticidade.
          </p>
        </div>

        {/* Coluna DIREITA: foto */}
        <div className="relative">
          <div className="absolute -inset-4 rounded-3xl bg-white/5 blur-2xl" />
          <div className="relative overflow-hidden rounded-3xl border border-white/10 bg-white/5">
            <Image
              src="/images/fotojustke.png"    // <-- AQUI VAI SUA FOTO (public/images/justke.jpg)
              alt="Foto de Richardt Justke"
              width={1200}
              height={1400}
              className="h-auto w-full object-cover"
              priority
            />
          </div>
        </div>
      </div>
    </section>
  );
}
