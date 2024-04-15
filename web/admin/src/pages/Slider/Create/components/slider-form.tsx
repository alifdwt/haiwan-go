import { Button } from "@/components/ui/button";
import { Heading } from "@/components/ui/heading";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import Slider from "@/interface/Slider";
import useCreateSlider from "./hooks/CreateSlider";
import { TrashIcon } from "lucide-react";
import useDeleteSlider from "./hooks/DeleteSlider";
import { AlertModal } from "@/components/modals/alert-modals";

interface SliderFormProps {
  initialData: Slider | null;
}

export const SliderForm: React.FC<SliderFormProps> = ({ initialData }) => {
  const {
    // open,
    // setOpen,
    name,
    file,
    setFile,
    handleNameChange,
    handleSubmit,
    handleFileUpload,
  } = useCreateSlider();

  const { handleDeleteSlider, isLoading, open, setOpen } = useDeleteSlider(
    initialData?.ID as number
  );

  const title = initialData ? "Edit Slider" : "Create Slider";
  const description = initialData ? "Edit slider data" : "Create new slider";
  // const toastMessage = initialData ? "Slider updated successfully" : "Slider created successfully"
  const action = initialData ? "Save Changes" : "Create";

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={handleDeleteSlider}
        loading={isLoading}
      />
      <div className="flex items-center justify-between">
        <Heading title={title} description={description} />
        {initialData && (
          <Button
            disabled={isLoading}
            variant="destructive"
            onClick={() => setOpen(true)}
          >
            <TrashIcon />
          </Button>
        )}
      </div>
      <Separator />
      <div className="space-y-8 w-full">
        {file && (
          <div className="flex">
            <div className="w-fit relative rounded-md">
              <div className="z-10 absolute top-2 right-2">
                <Button
                  type="button"
                  variant="destructive"
                  size={"icon"}
                  onClick={() => setFile(null)}
                >
                  <TrashIcon />
                </Button>
              </div>
              <img
                src={
                  initialData ? initialData.image : URL.createObjectURL(file)
                }
                alt="Preview"
                className="object-cover rounded-md"
              />
            </div>
          </div>
        )}
        <Button onClick={handleFileUpload} variant="outline">
          Upload Image
        </Button>
        <div className="grid grid-cols-3 gap-4">
          <Input
            placeholder="Name"
            value={initialData ? initialData.name : name}
            onChange={handleNameChange}
          />
        </div>
        <Button className="ml-auto" type="submit" onClick={handleSubmit}>
          {action}
        </Button>
      </div>
    </>
  );
};
