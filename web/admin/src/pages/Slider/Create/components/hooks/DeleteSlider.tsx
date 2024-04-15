import { toast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

export default function useDeleteSlider(sliderId: number) {
  const [isLoading, setIsLoading] = useState(false);
  const [open, setOpen] = useState(false);

  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { access_token } = useSelector((state: RootState) => state.user);

  const mutation = useMutation({
    mutationFn: () => {
      return myApi.delete(`/slider/delete/${sliderId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
    },
    onSuccess: () => {
      toast({
        title: "Please wait",
        description: "Slider is being deleted",
      });
      setIsLoading(true);
      setTimeout(() => {
        toast({
          title: "Slider deleted",
          description: "Slider deleted successfully",
        });
        queryClient.invalidateQueries({ queryKey: ["sliders"] });
        setIsLoading(false);
        setOpen(false);
      }, 3000);
      navigate("/sliders");
    },
    onError: () => {
      toast({
        title: "Failed to delete slider",
        description: "Something went wrong",
        variant: "destructive",
      });
    },
  });

  const handleDeleteSlider = () => {
    // e.stopPropagation();
    mutation.mutate();
  };

  return { handleDeleteSlider, isLoading, open, setOpen };
}
