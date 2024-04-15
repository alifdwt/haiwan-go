import { useToast } from "@/components/ui/use-toast";
import { SliderColumns } from "./columns";
import { AlertModal } from "@/components/modals/alert-modals";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuTrigger,
  DropdownMenuItem,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { CopyIcon, EditIcon, MoreHorizontal, TrashIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";
import useDeleteSlider from "../Create/components/hooks/DeleteSlider";

interface CellActionProps {
  data: SliderColumns;
}

export const CellAction: React.FC<CellActionProps> = ({ data }) => {
  const { toast } = useToast();
  const navigate = useNavigate();

  const onCopy = (id: number) => {
    navigator.clipboard.writeText(id.toString());
    toast({
      title: "Tersalin",
      description: "ID dari Slider berhasil disalin",
    });
  };

  const { handleDeleteSlider, isLoading, open, setOpen } = useDeleteSlider(
    data.ID
  );

  return (
    <>
      <AlertModal
        isOpen={open}
        onClose={() => setOpen(false)}
        onConfirm={handleDeleteSlider}
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
          <DropdownMenuLabel>Actions</DropdownMenuLabel>
          <DropdownMenuItem
            onClick={() => onCopy(data.ID)}
            className="cursor-pointer"
          >
            <CopyIcon className="mr-2 h-4 w-4" />
            Copy Id
          </DropdownMenuItem>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={() => navigate(`/sliders/${data.ID}`)}
          >
            <EditIcon className="mr-2 h-4 w-4" />
            Update
          </DropdownMenuItem>
          <DropdownMenuItem
            className="cursor-pointer"
            onClick={() => setOpen(true)}
          >
            <TrashIcon className="mr-2 h-4 w-4" />
            Delete
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
};
