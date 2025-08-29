import Header from "@/components/Header";

export default function PortfolioLayout({ children }) {
  return (
    <>
      <Header />
      {children}
    </>
  );
}
