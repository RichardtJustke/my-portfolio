"use client";
import { useEffect, useState } from "react";

/**
 * words: array de frases a digitar
 * typingSpeed: ms por caractere ao digitar
 * pauseTime: ms de pausa quando a palavra termina
 * loop: se true, repete; se false, digita uma vez e para
 */
export default function TypingAnimation({
  words = [],
  typingSpeed = 80,
  pauseTime = 1200,
  loop = true,
  className = "",
  cursorClassName = "inline-block w-[1ch] animate-pulse",
}) {
  const [i, setI] = useState(0);          // qual palavra
  const [sub, setSub] = useState(0);      // quantos chars
  const [del, setDel] = useState(false);  // deletando?

  useEffect(() => {
    if (!words.length) return;
    const current = words[i % words.length];

    // decide próximo passo
    let t;
    if (!del && sub < current.length) {
      t = setTimeout(() => setSub(sub + 1), typingSpeed);
    } else if (!del && sub === current.length) {
      if (loop) t = setTimeout(() => setDel(true), pauseTime);
    } else if (del && sub > 0) {
      t = setTimeout(() => setSub(sub - 1), typingSpeed / 2);
    } else {
      setDel(false);
      setI((i + 1) % words.length);
    }

    return () => clearTimeout(t);
  }, [sub, del, i, words, typingSpeed, pauseTime, loop]);

  const text = words.length ? words[i % words.length].slice(0, sub) : "";

  return (
    <span className={className}>
      {text}
      <span className={cursorClassName}>|</span>
    </span>
  );
}
