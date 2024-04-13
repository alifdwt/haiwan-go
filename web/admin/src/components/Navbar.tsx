import { PawPrint } from "lucide-react";
import Container from "./ui/container";
import { MainNav } from "./navbar/MainNav";
import NavbarAction from "./navbar/NavbarAction";

const Navbar = () => {
  return (
    <div className="bg-primary">
      <Container>
        <div className="relative px-4 sm:px-6 lg:px-8 flex h-16 items-center">
          <a href="/" className="ml-4 flex lg:ml-0 gap-x-2 items-center">
            <PawPrint size={24} className="text-secondary" />
            <p className="font-bold text-xl text-white">Haiwan</p>
          </a>
          <MainNav className="mx-4" />
          <NavbarAction />
        </div>
      </Container>
    </div>
  );
};

export { Navbar };
