import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { Product } from "@/interface/Product";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ChangeEvent, useCallback, useState } from "react";
import { useSelector } from "react-redux";

interface CreateCartValues {
  image_product: string;
  name: string;
  price: string;
  product_id: number;
  quantity: number;
  weight: number;
}

export default function useCart(product: Product) {
  const { access_token } = useSelector((state: RootState) => state.user);
  const queryClient = useQueryClient();
  const { toast } = useToast();

  const [quantity, setQuantity] = useState(1);

  const mutation = useMutation({
    mutationFn: (newCart: CreateCartValues) =>
      myApi.post("/cart/create", newCart, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      }),
    onSuccess: () => {
      toast({
        title: "Produk ditambahkan ke keranjang",
        description: "Produk ditambahkan ke keranjang",
      });
      queryClient.invalidateQueries({ queryKey: ["cart"] });
    },
    onError: (error) => {
      console.log(error);
      toast({
        title: "Gagal menambahkan produk ke keranjang",
        description: error.message,
      });
    },
  });

  const handleQuantityChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      setQuantity(Number(e.target.value));
    },
    []
  );

  const handleAddToCart = useCallback(() => {
    mutation.mutate({
      image_product: product.image_path,
      name: product.name,
      price: String(product.price),
      product_id: product.id,
      quantity: quantity,
      weight: product.weight,
    });
  }, [mutation, product, quantity]);

  return {
    quantity,
    handleQuantityChange,
    handleAddToCart,
  };
}
