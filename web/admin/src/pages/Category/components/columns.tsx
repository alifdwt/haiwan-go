import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type CategoryColumns = {
  id: number;
  name: string;
  description: string;
  slug: string;
  image_path: string;
  created_at: string;
  updated_at: string;
};

export const columns: ColumnDef<CategoryColumns>[] = [
  {
    header: "No",
    cell: (info) => info.row.index + 1,
  },
  {
    accessorKey: "image_path",
    header: "Image",
    cell: ({ row }) => {
      return (
        <img
          src={row.original.image_path}
          alt={row.original.name}
          className="h-10 w-10"
        />
      );
    },
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "description",
    header: "Description",
  },
  {
    id: "actions",
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
