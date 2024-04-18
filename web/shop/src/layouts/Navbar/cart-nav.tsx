import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ShoppingCartIcon } from "lucide-react";

export default function CartNav() {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <Button variant={"secondary"}>
          <ShoppingCartIcon className="h-4 w-4" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-96" align="end">
        <DropdownMenuLabel className="font-normal flex justify-between">
          <p>Keranjang</p>
          <a href="/cart">Lihat Keranjang</a>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
