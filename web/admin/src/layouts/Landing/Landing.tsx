import { Navbar } from "@/components/Navbar";
import { Outlet } from "react-router-dom";

export default function LandingLayout() {
  return (
    <div className="min-h-screen">
      <Navbar />
      <main>
        <Outlet />
      </main>
    </div>
  );
}
