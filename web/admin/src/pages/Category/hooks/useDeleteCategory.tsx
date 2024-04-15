import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

export default function useDeleteCategory(categoryId: number) {
  const { toast } = useToast();
  const [isLoading, setIsLoading] = useState(false);
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { access_token } = useSelector((state: RootState) => state.user);

  const mutation = useMutation({
    mutationFn: () => {
      return myApi.delete(`/category/delete/${categoryId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
    },
    onSuccess: () => {
      toast({
        title: "Please wait",
        description: "Category is being deleted",
      });
      setIsLoading(true);
      setTimeout(() => {
        toast({
          title: "Category deleted",
          description: "Category deleted successfully",
        });
        queryClient.invalidateQueries({ queryKey: ["categories"] });
        setIsLoading(false);
        setOpen(false);
      }, 3000);
      navigate("/categories");
    },
    onError: (error) => {
      toast({
        title: "Failed to delete category",
        description: error.message,
        variant: "destructive",
      });
    },
  });

  const handleDeleteCategory = () => {
    mutation.mutate();
  };

  return { handleDeleteCategory, isLoading, open, setOpen };
}
