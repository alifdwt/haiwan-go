import { PawPrintIcon } from "lucide-react";
import MainNav from "./MainNavbar";
import LogOutButton from "./components/LogoutButton";

export default function Navbar() {
  return (
    <div className="border-b">
      <div className="flex h-16 items-center px-4">
        <PawPrintIcon size={24} className="text-secondary" />
        <MainNav className="mx-4" />
        <LogOutButton />
      </div>
    </div>
  );
}
