import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";

export default function useCreateCategory() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { access_token } = useSelector((state: RootState) => state.user);

  const [createIsLoading, setCreateIsLoading] = useState(false);
  const [open, setOpen] = useState(false);

  const [name, setName] = useState("");
  // const [description, setDescription] = useState("");

  const [file, setFile] = useState<File | null>(null);

  const { toast } = useToast();
  const form = new FormData();

  const mutation = useMutation({
    mutationFn: (newCategory: FormData) =>
      myApi.post("/category/create", newCategory, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      }),
    onSuccess: () => {
      toast({
        title: "Category is being created",
        description: "Category is being created, please wait",
      });
      setCreateIsLoading(true);
      setTimeout(() => {
        toast({
          title: "Category created",
          description: "Category created successfully",
        });
        queryClient.invalidateQueries({ queryKey: ["categories"] });
        setCreateIsLoading(false);
        setOpen(false);
      }, 3000);
      setName("");
      setFile(null);
      navigate("/categories");
    },
    onError: () => {
      toast({
        title: "Category failed to be created",
        description: "Please try again",
      });
      setCreateIsLoading(false);
      setOpen(false);
    },
  });

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleSubmit = () => {
    form.append("name", name);
    form.append("image_category", file as File);

    mutation.mutate(form);
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
    input.id = "image_category";
    input.name = "image_category";

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
