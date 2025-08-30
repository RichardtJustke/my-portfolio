"use client";
import { cn } from "@/lib/utils";
import { IconMenu2, IconX } from "@tabler/icons-react";
import {
  motion,
  AnimatePresence,
  useScroll,
  useMotionValueEvent,
} from "motion/react";
import React, { useRef, useState } from "react";

/** Nav “wrapper” que muda largura/estilo conforme scroll */
export const Navbar = ({ children, className }) => {
  const ref = useRef(null);
  const { scrollY } = useScroll({ target: ref, offset: ["start start", "end start"] });
  const [visible, setVisible] = useState(false);

  useMotionValueEvent(scrollY, "change", (latest) => {
    setVisible(latest > 100);
  });

  return (
    <motion.div
      ref={ref}
      // use "fixed top-0" se quiser fixo; "sticky top-0" se quiser grudado no fluxo
      className={cn("fixed inset-x-0 top-0 z-50 w-full", className)}
    >
      {React.Children.map(children, (child) =>
        React.isValidElement(child)
          ? React.cloneElement(child, { visible })
          : child
      )}
    </motion.div>
  );
};

export const NavBody = ({ children, className, visible }) => {
  return (
    <motion.div
      animate={{
        backdropFilter: visible ? "blur(10px)" : "none",
        boxShadow: visible
          ? "0 0 24px rgba(34,42,53,.06), 0 1px 1px rgba(0,0,0,.05), 0 0 0 1px rgba(34,42,53,.04), 0 0 4px rgba(34,42,53,.08), 0 16px 68px rgba(47,48,55,.05), 0 1px 0 rgba(255,255,255,.1) inset"
          : "none",
        width: visible ? "60%" : "100%", // ajustei pra caber no teu layout
        y: visible ? 10 : 0,
      }}
      transition={{ type: "spring", stiffness: 200, damping: 50 }}
      style={{ minWidth: "280px" }}
      className={cn(
        "mx-auto flex w-full max-w-7xl items-center justify-between rounded-full bg-transparent px-4 py-3",
        "border-b border-white/10 bg-black/60 backdrop-blur-md", // mantém teu visual
        className
      )}
    >
      {children}
    </motion.div>
  );
};

export const NavItems = ({ items = [], className, onItemClick }) => {
  const [hovered, setHovered] = useState(null);
  return (
    <motion.nav
      onMouseLeave={() => setHovered(null)}
      className={cn(
        "hidden lg:flex flex-1 items-center justify-center gap-2 text-sm font-medium",
        className
      )}
    >
      {items.map((item, idx) => (
        <a
          key={item.link}
          href={item.link}
          onMouseEnter={() => setHovered(idx)}
          onClick={onItemClick}
          className="relative px-4 py-2 text-zinc-300 hover:text-white"
        >
          {hovered === idx && (
            <motion.span
              layoutId="hovered-pill"
              className="absolute inset-0 rounded-full bg-white/10"
            />
          )}
          <span className="relative z-10">{item.name}</span>
        </a>
      ))}
    </motion.nav>
  );
};

export const MobileNav = ({ children, className, visible }) => {
  return (
    <motion.div
      animate={{
        backdropFilter: visible ? "blur(10px)" : "none",
        y: visible ? 6 : 0,
      }}
      transition={{ type: "spring", stiffness: 200, damping: 50 }}
      className={cn(
        "lg:hidden mx-auto w-full max-w-[calc(100vw-2rem)] px-2 py-2",
        className
      )}
    >
      {children}
    </motion.div>
  );
};

export const MobileNavHeader = ({ children, className }) => (
  <div className={cn("flex items-center justify-between w-full", className)}>{children}</div>
);

export const MobileNavMenu = ({ children, className, isOpen, onClose }) => (
  <AnimatePresence>
    {isOpen && (
      <motion.div
        initial={{ opacity: 0, y: -6 }}
        animate={{ opacity: 1, y: 0 }}
        exit={{ opacity: 0, y: -6 }}
        className={cn(
          "absolute inset-x-2 top-14 z-[60] rounded-xl bg-black/90 p-4 border border-white/10",
          className
        )}
      >
        {children}
      </motion.div>
    )}
  </AnimatePresence>
);

export const MobileNavToggle = ({ isOpen, onClick }) =>
  isOpen ? (
    <IconX className="text-white" onClick={onClick} />
  ) : (
    <IconMenu2 className="text-white" onClick={onClick} />
  );
