import { AlertModal } from "@/components/modals/alert-modals";
import { Button } from "@/components/ui/button";
import { Heading } from "@/components/ui/heading";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { myApi } from "@/helpers/api";
import Slider from "@/interface/Slider";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import * as z from "zod";

const formSchema = z.object({
  name: z.string().min(1, { message: "Name is required" }),
  file: z.instanceof(File),
});

type SliderFormValues = z.infer<typeof formSchema>;

interface SliderFormProps {
  initialData: Slider | null;
}

export const SliderForm: React.FC<SliderFormProps> = ({ initialData }) => {
  const [open, setOpen] = useState(false);

  const title = initialData ? "Edit Slider" : "Create Slider";
  const description = initialData ? "Edit slider data" : "Create new slider";
  // const toastMessage = initialData ? "Slider updated successfully" : "Slider created successfully"
  const action = initialData ? "Save Changes" : "Create";

  const mutation = useMutation({
    mutationFn: (data: SliderFormValues) => {
      return myApi.post("/slider", data);
    },
  });

  const onSubmit = () => {
    if (initialData) {
      // update
    } else {
      // create
    }
  };

  const onDelete = () => {
    // delete
  };

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={onDelete}
        loading={mutation.isPending}
      />
      <div className="flex items-center justify-between">
        <Heading title={title} description={description} />
        {initialData && (
          <Button
            disabled={mutation.isPending}
            variant="destructive"
            onClick={() => setOpen(true)}
          >
            Delete
          </Button>
        )}
      </div>
      <Separator />
      <div className="space-y-8 w-full">
        <Input type="file" />
        <div className="grid grid-cols-3 gap-4">
          <Input placeholder="Name" />
        </div>
        <Button className="ml-auto" type="submit" onClick={onSubmit}>
          {action}
        </Button>
      </div>
    </>
  );
};
