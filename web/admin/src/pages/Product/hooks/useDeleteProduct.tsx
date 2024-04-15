import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

export default function useDeleteProduct(productId: number) {
  const { toast } = useToast();
  const [isLoading, setIsLoading] = useState(false);
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { access_token } = useSelector((state: RootState) => state.user);

  const mutation = useMutation({
    mutationFn: () => {
      return myApi.delete(`/product/delete/${productId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
    },
    onSuccess: () => {
      toast({
        title: "Please wait",
        description: "Product is being deleted",
      });
      setIsLoading(true);
      setTimeout(() => {
        toast({
          title: "Product deleted",
          description: "Product deleted successfully",
        });
        queryClient.invalidateQueries({ queryKey: ["products"] });
        setIsLoading(false);
        setOpen(false);
      }, 3000);
      navigate("/products");
    },
    onError: (error) => {
      toast({
        title: "Product failed to be deleted",
        description: `${error.message}`,
      });
      setIsLoading(false);
      setOpen(false);
    },
  });

  const handleDeleteProduct = () => {
    mutation.mutate();
  };

  return { handleDeleteProduct, isLoading, open, setOpen };
}
