import { useToast } from "@/components/ui/use-toast";
import { CategoryColumns } from "./columns";
import { useNavigate } from "react-router-dom";
import { AlertModal } from "@/components/modals/alert-modals";
import useDeleteCategory from "../hooks/useDeleteCategory";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { CopyIcon, EditIcon, MoreHorizontal, TrashIcon } from "lucide-react";

interface CellActionProps {
  data: CategoryColumns;
}

export const CellAction: React.FC<CellActionProps> = ({ data }) => {
  const { toast } = useToast();
  const navigate = useNavigate();

  const onCopy = (id: number) => {
    navigator.clipboard.writeText(id.toString());
    toast({
      title: "Tersalin",
      description: "ID dari Kategori berhasil disalin",
    });
  };

  const { handleDeleteCategory, isLoading, open, setOpen } = useDeleteCategory(
    data.id
  );

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={handleDeleteCategory}
        loading={isLoading}
      />
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" className="w-8 h-8 p-0">
            <span className="sr-only">Open menu</span>
            <MoreHorizontal className="h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <DropdownMenuLabel>Fungsi</DropdownMenuLabel>
          <DropdownMenuItem
            onClick={() => onCopy(data.id)}
            className="cursor-pointer"
          >
            <CopyIcon className="mr-2 h-4 w-4" />
            Salin ID
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={() => navigate(`/categories/${data.id}`)}
            className="cursor-pointer"
          >
            <EditIcon className="mr-2 h-4 w-4" />
            Ubah
          </DropdownMenuItem>
          <DropdownMenuItem
            onClick={() => setOpen(true)}
            className="cursor-pointer"
          >
            <TrashIcon className="mr-2 h-4 w-4" />
            Hapus
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
};
