import { Button } from "@/components/ui/button";
import { ArrowRightFromLine, CheckCircle, Leaf, Store } from "lucide-react";

const perks = [
  {
    name: "Berjualan dengan Mudah",
    Icon: Store,
    description: "Jual beli peralatan hewan tanpa harus keluar rumah.",
  },
  {
    name: "Tersedia Setiap Saat",
    Icon: CheckCircle,
    description: "Kami akan melayani anda setiap saat.",
  },
  {
    name: "Hijaukan Dunia",
    Icon: Leaf,
    description: "Kami berkomitmen untuk membuat dunia lebih hijau.",
  },
];

export default function Home() {
  return (
    <>
      <div className="py-20 bg-daintree-950">
        <div className="text-center flex flex-col items-center mx-auto max-w-3xl">
          <h3 className="px-5 py-3 bg-daintree-900 rounded-full text-white uppercase">
            Haiwan Web Admin
          </h3>
          <h1 className="mt-6 text-4xl font-bold sm:text-6xl uppercase text-white">
            Tempat Jual Beli{" "}
            <span className="text-secondary">peralatan hewan</span>
          </h1>
          <p className="mt-6 text-lg max-w-prose text-muted-foreground">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 mt-6">
            <Button
              asChild
              variant={"secondary"}
              className="rounded-full px-5 hover:bg-honeysuckle-500"
            >
              <a href="/register" className="flex gap-2">
                Mulai Jualan
              </a>
            </Button>
            <Button
              asChild
              className="bg-daintree-200 text-primary rounded-full px-5 hover:bg-daintree-300"
            >
              <a href="/login" className="flex gap-2">
                <p>Lanjut Jualan</p> <ArrowRightFromLine size={18} />
              </a>
            </Button>
            {/* <Button
              asChild
              className="bg-daintree-200 text-primary rounded-full px-5 hover:bg-daintree-300"
            >
              <a href="#" className="flex gap-2">
                <p>Hubungi Kami</p> <ArrowRightFromLine size={18} />
              </a>
            </Button> */}
          </div>
        </div>
      </div>

      <section className="border-t border-gray-200 bg-gray-50">
        <div className="mx-auto w-full max-w-screen-xl px-2.5 md:px-20 py-20">
          <div className="grid grid-cols-1 gap-y-12 sm:grid-cols-2 sm:gap-x-6 lg:grid-cols-3 lg:gap-x-8 lg:gap-y-0">
            {perks.map((perk) => (
              <div
                key={perk.name}
                className="text-center md:flex md:items-start md:text-left lg:block lg:text-center"
              >
                <div className="md:flex-shrink-0 flex justify-center">
                  <div className="h-16 w-16 flex items-center justify-center rounded-full bg-daintree-200 text-daintree-800">
                    {<perk.Icon className="w-1/3 h-1/3" />}
                  </div>
                </div>

                <div className="mt-6 md:ml-4 md:mt-0 lg:ml-0 lg:mt-6">
                  <h3 className="text-base font-medium text-gray-900">
                    {perk.name}
                  </h3>
                  <p className="mt-3 text-sm text-muted-foreground">
                    {perk.description}
                  </p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>
    </>
  );
}
