import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useSelector } from "react-redux";

export default function useDeleteCart() {
  const { access_token } = useSelector((state: RootState) => state.user);
  const queryClient = useQueryClient();
  const { toast } = useToast();

  const mutation = useMutation({
    mutationFn: (cartId: number) => {
      return myApi.delete(`/cart/delete/${cartId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
    },
    onSuccess: () => {
      toast({
        title: "Produk di hapus dari keranjang",
        description: "Produk di hapus dari keranjang",
      });
      queryClient.invalidateQueries({ queryKey: ["cart"] });
    },
    onError: (error) => {
      console.log(error);
      toast({
        title: "Gagal menghapus produk dari keranjang",
        description: error.message,
      });
    },
  });

  const handleDeleteCart = (cartId: number) => {
    mutation.mutate(cartId);
  };

  return { handleDeleteCart };
}
