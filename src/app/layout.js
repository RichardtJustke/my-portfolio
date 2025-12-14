import "./globals.css";

export const metadata = {
  title: "Ketsu",
  description: "colocar o nome da pagina",
  icons: {
    icon: "/simbloColor.svg",
  },
};

export default function RootLayout({ children }) {
  return (
    <html lang="pt-BR">
      <body>{children}</body>
    </html>
  );
}
