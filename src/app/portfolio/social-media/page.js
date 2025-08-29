export const metadata = { title: "Social Media & Conteúdo – Portfólio" };

export default function SocialMedia() {
  return (
    <main className="min-h-screen bg-black text-white px-6 py-20">
      <div className="max-w-5xl mx-auto">
        <a href="/portfolio" className="opacity-70 hover:opacity-100">← Voltar ao Portfólio</a>
        <h1 className="mt-4 text-4xl md:text-6xl font-extrabold">Social Media</h1>
        <p className="mt-3 opacity-80">Posts, carrosséis, campanhas e motion.</p>
      </div>
    </main>
  );
}
