import { Outlet } from "react-router-dom";
import Navbar from "./Navbar/Navbar";
import Footer from "./Footer/Footer";

export default function Layout() {
  return (
    <div className="">
      <Navbar />
      <main className="w-[80vw] mx-auto mb-4">
        <Outlet />
      </main>
      <Footer />
    </div>
  );
}
