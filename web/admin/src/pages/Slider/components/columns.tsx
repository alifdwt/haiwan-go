import { Button } from "@/components/ui/button";
import { ColumnDef } from "@tanstack/react-table";

export type SliderColumns = {
  id: number;
  name: string;
  created_at: string;
};

export const columns: ColumnDef<SliderColumns>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "created_at",
    header: "Created At",
  },
  {
    id: "actions",
    cell: () => <Button>Edit</Button>,
  },
];
