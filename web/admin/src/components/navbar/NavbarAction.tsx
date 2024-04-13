import { ArrowRightFromLine } from "lucide-react";
import { Button } from "../ui/button";

const NavbarAction = () => {
  return (
    <div className="ml-auto items-center gap-x-4">
      <Button
        variant={"secondary"}
        className="flex items-center rounded-full px-4 py-2"
      >
        <ArrowRightFromLine className="mr-2 h-4 w-4" />
        Mulai Jualan
      </Button>
    </div>
  );
};

export default NavbarAction;
