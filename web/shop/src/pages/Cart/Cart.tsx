import TransactionBreadcrumb from "@/components/transaction-breadcrumb";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { myApi } from "@/helpers/api";
import { ToRupiah } from "@/lib/currency";
import { getTotalPrice, getTotalQuantity } from "@/lib/price";
import { getCarts } from "@/slices/cart/cartSlice";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { Trash2Icon } from "lucide-react";
import { useDispatch, useSelector } from "react-redux";

export default function CartPage() {
  const dispatch = useDispatch();
  const { access_token } = useSelector((state: RootState) => state.user);
  const { carts } = useSelector((state: RootState) => state.cart);

  const { isLoading } = useQuery({
    queryKey: ["cart"],
    queryFn: async () => {
      const { data } = await myApi.get("/cart", {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
      dispatch(getCarts(data.data));
      return data.data;
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="max-w-screen-xl mx-auto">
      <TransactionBreadcrumb />
      <div className="my-5 grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="col-span-2 flex flex-col gap-2">
          {carts.map((cart) => (
            <div key={cart.id} className="flex flex-col gap-2 border p-4">
              <div className="flex justify-between gap-4 w-full">
                <div className="flex items-center space-x-3">
                  <img
                    src={cart.image}
                    alt={cart.name}
                    className="w-12 h-12 object-cover rounded-lg"
                  />
                  <div className="flex flex-col ml-4">
                    <p>{cart.name}</p>
                    <p>
                      {cart.quantity} x {ToRupiah(parseInt(cart.price))}
                    </p>
                  </div>
                </div>
                <Button variant={"destructive"} size={"icon"}>
                  <Trash2Icon className="w-5 h-5" />
                </Button>
              </div>
              <Separator />
              <div className="flex justify-between gap-4 w-full">
                <Input type="text" placeholder="Tulis catatan" />
                <div className="flex gap-2 items-center">
                  <Button
                    variant={"secondary"}
                    className="flex gap-2 rounded-full"
                  >
                    -
                  </Button>
                  <div>
                    <span>{cart.quantity}</span>
                  </div>
                  <Button
                    variant={"secondary"}
                    className="flex gap-2 rounded-full"
                  >
                    +
                  </Button>
                </div>
              </div>
            </div>
          ))}
        </div>
        <Card>
          <CardHeader className="space-y-2">
            <CardTitle>Rangkuman</CardTitle>
            <div className="flex gap-2">
              <Input type="text" placeholder="Anda punya kode promo" />
              <Button variant={"secondary"}>Cek</Button>
            </div>
          </CardHeader>
          <CardContent className="space-y-2">
            <div className="flex justify-between">
              <p>Total Harga ({getTotalQuantity(carts)} item)</p>
              <p className="font-bold">{ToRupiah(getTotalPrice(carts))}</p>
            </div>
            <Separator />
            <div className="flex justify-between">
              <p>Total Harga ({getTotalQuantity(carts)} item)</p>
              <p className="font-bold">{ToRupiah(getTotalPrice(carts))}</p>
            </div>
          </CardContent>
          <CardFooter>
            <Button className="w-full" asChild>
              <a href="/cart/shipment">Lanjutkan</a>
            </Button>
          </CardFooter>
        </Card>
      </div>
    </div>
  );
}
