import { Outlet } from "react-router-dom";
import Navbar from "./Navbar/Navbar";

export default function Layout() {
  return (
    <div className="min-h-screen">
      <Navbar />
      <main className="p-4">
        <Outlet />
      </main>
    </div>
  );
}
