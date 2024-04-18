import { Outlet } from "react-router-dom";
import Navbar from "./Navbar/Navbar";
import Footer from "./Footer/Footer";

export default function Layout() {
  return (
    <div className="min-h-screen flex flex-col justify-between">
      <Navbar />
      <main className="w-[80vw] mx-auto mb-4">
        <Outlet />
      </main>
      <Footer />
    </div>
  );
}
