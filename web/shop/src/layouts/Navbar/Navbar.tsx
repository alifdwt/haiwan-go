import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { PawPrintIcon } from "lucide-react";

export default function Navbar() {
  return (
    <div className="border-b bg-primary">
      <div className="flex h-16 items-center px-4 gap-4">
        <div className="flex gap-x-2 items-center">
          <PawPrintIcon size={24} className="text-secondary" />
          <p className="font-bold text-xl text-white">Haiwan</p>
        </div>
        <div className="flex gap-x-2 items-center">
          <Input className="w-72" placeholder="Search..." />
          <Button variant={"secondary"}>Search</Button>
        </div>
        <Button variant={"secondary"} className="ml-auto">
          Alif Dewantara
        </Button>
      </div>
    </div>
  );
}
