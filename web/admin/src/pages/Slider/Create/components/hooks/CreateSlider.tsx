import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

export default function useCreateSlider() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { access_token } = useSelector((state: RootState) => state.user);

  const [createIsLoading, setCreateIsLoading] = useState(false);
  const [open, setOpen] = useState(false);
  const [name, setName] = useState("");
  const [file, setFile] = useState<File | null>(null);

  const { toast } = useToast();
  const form = new FormData();

  const mutation = useMutation({
    mutationFn: (newSlider: FormData) =>
      myApi.post("/slider/create", newSlider, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      }),
    onSuccess: () => {
      toast({
        title: "Slider is being created",
        description: "Slider is being created, please wait",
      });
      setCreateIsLoading(true);
      setTimeout(() => {
        toast({
          title: "Slider created",
          description: "Slider created successfully",
        });
        queryClient.invalidateQueries({ queryKey: ["sliders"] });
        setCreateIsLoading(false);
      }, 5000);
      navigate("/sliders");
    },
    onError: () => {
      toast({
        title: "Failed to create slider",
        description: "Something went wrong",
        variant: "destructive",
      });
    },
  });

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleSubmit = () => {
    form.append("name", name);
    form.append("file", file as File);

    mutation.mutate(form);
    setName("");
    setFile(null);
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setFile(file);
    }
  };

  const handleFileUpload = () => {
    const input = document.createElement("input");
    input.type = "file";
    input.id = "file";
    input.name = "file";

    // @ts-expect-error next-line
    input.onchange = handleFileChange;
    input.click();
  };

  return {
    open,
    setOpen,
    name,
    file,
    setFile,
    handleNameChange,
    handleSubmit,
    handleFileUpload,
    createIsLoading,
  };
}
