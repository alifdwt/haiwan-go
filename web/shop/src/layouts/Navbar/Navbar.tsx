import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { PawPrintIcon } from "lucide-react";
import CartNav from "./cart-nav";
import UserNav from "./user-nav";

export default function Navbar() {
  return (
    <div className="border-b bg-primary">
      <div className="flex h-16 items-center px-4 gap-4 justify-between">
        <a className="flex gap-x-2 items-center" href="/">
          <PawPrintIcon size={24} className="text-secondary" />
          <p className="font-bold text-xl text-white">Haiwan</p>
        </a>
        <div className="flex gap-x-2 items-center">
          <Input className="w-72" placeholder="Search..." />
          <Button variant={"secondary"}>Search</Button>
        </div>
        <div className="flex gap-x-2 items-center">
          <CartNav />
          <UserNav />
        </div>
      </div>
    </div>
  );
}
