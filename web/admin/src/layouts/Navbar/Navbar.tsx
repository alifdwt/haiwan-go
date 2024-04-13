import { PawPrintIcon } from "lucide-react";
import MainNav from "./MainNavbar";
import LogOutButton from "./components/LogoutButton";

export default function Navbar() {
  return (
    <div className="border-b bg-primary">
      <div className="flex h-16 items-center px-4">
        <div className="flex gap-x-2 items-center">
          <PawPrintIcon size={24} className="text-secondary" />
          <p className="font-bold text-xl text-white">Haiwan</p>
        </div>
        <MainNav className="mx-4" />
        <LogOutButton />
      </div>
    </div>
  );
}
