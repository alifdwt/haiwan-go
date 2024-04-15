import { CategoryColumns } from "@/pages/Category/components/columns";
import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./cell-action";

export type ProductColumns = {
  id: number;
  name: string;
  description: string;
  price: number;
  count_in_stock: string;
  brand: string;
  weight: number;
  rating: string;
  image_path: string;
  created_at: string;
  category: CategoryColumns;
};

export const columns: ColumnDef<ProductColumns>[] = [
  {
    header: "No",
    cell: (info) => info.row.index + 1,
  },
  {
    accessorKey: "image_path",
    header: "Image",
    cell: ({ row }) => {
      return (
        <div className="w-16 h-16">
          <img src={row.original.image_path} alt={row.original.name} />
        </div>
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
    accessorKey: "price",
    header: "Price",
  },
  {
    accessorKey: "count_in_stock",
    header: "Stock",
  },
  {
    accessorKey: "rating",
    header: "Rating",
  },
  {
    accessorKey: "created_at",
    header: "Created At",
    cell: ({ row }) => {
      // const time = TimeConverter(row.original.created_at);
      return <p>{row.original.created_at}</p>;
    },
  },
  {
    id: "actions",
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
