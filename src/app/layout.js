import "./globals.css";

export const metadata = {
  title: "Meu Site",
  description: "Portfólio do Justke",
};

export default function RootLayout({ children }) {
  return (
    <html lang="pt-BR">
      <body>{children}</body>
    </html>
  );
}
