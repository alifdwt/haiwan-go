import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { RootState } from "@/store";
import { useSelector } from "react-redux";
import useCart from "./hooks/useCart";
import { Product } from "@/interface/Product";

export default function ProductTransaction({ product }: { product: Product }) {
  const { access_token } = useSelector((state: RootState) => state.user);
  const { quantity, handleQuantityChange, handleAddToCart } = useCart(product);

  return (
    <div className="flex gap-2">
      <div className="w-3/4 flex flex-col gap-2">
        {access_token && (
          <Button
            variant={"secondary"}
            className="w-full"
            onClick={handleAddToCart}
          >
            Masukkan ke Keranjang
          </Button>
        )}
        <Button className="w-full">Beli Sekarang</Button>
      </div>
      <div className="w-1/4">
        <p className="text-sm text-center mb-2">Jumlah</p>
        <Input
          type="number"
          placeholder="Jumlah"
          className="w-full"
          value={quantity}
          onChange={handleQuantityChange}
        />
      </div>
    </div>
  );
}
