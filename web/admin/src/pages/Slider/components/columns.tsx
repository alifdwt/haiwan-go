import { TimeConverter } from "@/lib/converter";
import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type SliderColumns = {
  ID: number;
  name: string;
  image: string;
  CreatedAt: string;
};

export const columns: ColumnDef<SliderColumns>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "image",
    header: "Image",
    cell: ({ row }) => {
      return (
        <img
          src={row.original.image}
          alt={row.original.name}
          className="h-12 object-cover"
        />
      );
    },
  },
  {
    accessorKey: "CreatedAt",
    header: "Created At",
    cell: ({ row }) => {
      const time = TimeConverter(row.original.CreatedAt);
      return <p>{time}</p>;
    },
  },
  {
    id: "actions",
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
