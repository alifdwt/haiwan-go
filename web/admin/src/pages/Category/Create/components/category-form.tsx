import { Category } from "@/interface/Category";
import useCreateCategory from "../../hooks/useCreateCategory";
import { AlertModal } from "@/components/modals/alert-modals";
import useDeleteCategory from "../../hooks/useDeleteCategory";
import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { TrashIcon } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { Input } from "@/components/ui/input";

interface CategoryFormProps {
  initialData: Category | null;
}

export const CategoryForm: React.FC<CategoryFormProps> = ({ initialData }) => {
  const {
    name,
    file,
    setFile,
    handleNameChange,
    handleSubmit,
    handleFileUpload,
    createIsLoading,
  } = useCreateCategory();

  const { handleDeleteCategory, isLoading, open, setOpen } = useDeleteCategory(
    initialData?.id as number
  );

  const title = initialData ? "Edit Category" : "Create Category";
  const description = initialData
    ? "Edit category data"
    : "Create new category";
  const action = initialData ? "Save Changes" : "Create";

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={handleDeleteCategory}
        loading={isLoading}
      />
      <div className="flex items-center justify-between">
        <Heading title={title} description={description} />
        <Button
          variant="destructive"
          onClick={() => setOpen(true)}
          disabled={createIsLoading}
        >
          <TrashIcon className="w-4 h-4 mr-2" />
          Delete
        </Button>
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
                  initialData
                    ? initialData.image_path
                    : URL.createObjectURL(file)
                }
                alt="Preview"
                className="object-cover rounded-md w-40 h-40"
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
