import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { myApi } from "@/helpers/api";
import useDeleteCart from "@/hooks/useDeleteCart";
import { ToRupiah } from "@/lib/currency";
import { getTotalPrice } from "@/lib/price";
import { getCarts } from "@/slices/cart/cartSlice";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { ShoppingCartIcon, Trash2Icon } from "lucide-react";
import { useDispatch, useSelector } from "react-redux";

export default function CartNav() {
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
    enabled: !!access_token,
  });

  const { handleDeleteCart } = useDeleteCart();

  if (!access_token) return null;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <Button variant={"ghost"} className="flex">
          <ShoppingCartIcon className="h-4 w-4" color="white" />
          {carts && carts.length > 0 && (
            <sup className="text-xs bg-red-500 text-primary-foreground rounded-full px-1">
              {carts.length}
            </sup>
          )}
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-96" align="end">
        <DropdownMenuLabel className="font-normal flex justify-between">
          <p>Keranjang</p>
          <a href="/cart">Lihat Keranjang</a>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          {isLoading ? (
            <p>Memuat...</p>
          ) : (
            <>
              {carts ? (
                carts.map((cart) => (
                  <DropdownMenuItem
                    key={cart.id}
                    className="flex justify-between"
                  >
                    <div className="flex items-center space-x-3">
                      <img
                        src={cart.image}
                        alt={cart.name}
                        width={50}
                        height={50}
                        className="h-12 w-12 object-cover rounded-lg"
                      />
                      <div className="flex flex-col ml-4 w-64">
                        <p>{cart.name}</p>
                        <p>
                          {cart.quantity} x {ToRupiah(parseInt(cart.price))}
                        </p>
                      </div>
                      <Button
                        variant={"destructive"}
                        size={"icon"}
                        onClick={() => handleDeleteCart(cart.id)}
                      >
                        <Trash2Icon className="h-5 w-5" />
                      </Button>
                    </div>
                  </DropdownMenuItem>
                ))
              ) : (
                <p className="text-center my-5">Belum ada barang nih</p>
              )}
              <DropdownMenuSeparator />
              <DropdownMenuLabel className="font-normal flex justify-between">
                <p>Total</p>
                <p className="font-bold">
                  {carts && ToRupiah(getTotalPrice(carts))}
                </p>
              </DropdownMenuLabel>
            </>
          )}
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
