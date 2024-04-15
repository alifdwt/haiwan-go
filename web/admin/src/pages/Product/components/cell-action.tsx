import { useToast } from "@/components/ui/use-toast";
import { ProductColumns } from "./columns";
import { useNavigate } from "react-router-dom";
import useDeleteProduct from "../hooks/useDeleteProduct";
import { AlertModal } from "@/components/modals/alert-modals";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import {
  CopyIcon,
  EditIcon,
  MoreHorizontalIcon,
  TrashIcon,
} from "lucide-react";

interface CellActionProps {
  data: ProductColumns;
}

export const CellAction: React.FC<CellActionProps> = ({ data }) => {
  const { toast } = useToast();
  const navigate = useNavigate();

  const onCopy = (id: number) => {
    navigator.clipboard.writeText(id.toString());
    toast({
      title: "Tersalin",
      description: "ID dari Produk berhasil disalin",
    });
  };

  const { handleDeleteProduct, isLoading, open, setOpen } = useDeleteProduct(
    data.id
  );

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={handleDeleteProduct}
        loading={isLoading}
      />
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" className="w-8 h-8 p-0">
            <span className="sr-only">Open menu</span>
            <MoreHorizontalIcon className="h-4 w-4" />
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
            onClick={() => navigate(`/products/${data.id}`)}
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
