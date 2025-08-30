import Image from "next/image";

const logos = [
  "/images/log1.png",
  "/images/log2.png",
  "/images/log3.png",
  "/images/log4.png",
];

function Row() {
  return (
    <>
      {logos.map((src, i) => (
        <Image
          key={i}
          src={src}
          alt={`Logo ${i + 1}`}
          width={160}
          height={80}
          className="h-20 w-auto object-contain"
        />
      ))}
    </>
  );
}

export default function InfiniteLogos() {
  return (
    <section className="mx-auto w-full max-w-6xl px-4 md:px-6 py-16">
      {/* Marquee */}
      <div className="relative overflow-hidden scroller-mask">
        <div
          className="marquee-track flex items-center gap-20"
          style={{ ["--speed"]: "28s" }}
        >
          <Row />
          <Row /> {/* duplicado pro looping infinito */}
        </div>
      </div>
    </section>
  );
}
