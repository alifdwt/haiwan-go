import { Heading } from "@/components/ui/heading";
import { CategoryColumns, columns } from "./columns";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { DataTable } from "@/components/ui/data-table";
import { ApiList } from "@/components/ui/api-list";

interface CategoryClientProps {
  data: CategoryColumns[];
}

export const CategoryClient: React.FC<CategoryClientProps> = ({ data }) => {
  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Kategori (${data.length})`}
          description="Lihat dan atur kategorimu"
        />
        <a href="/categories/create">
          <Button variant={"secondary"}>+ Tambah Kategori</Button>
        </a>
      </div>
      <Separator />
      <DataTable searchKey="name" columns={columns} data={data} />
      <Heading title="API" description="Dokumentasi API untuk kategori" />
      <Separator />
      <ApiList entityName="category" entityIdName="categoryId" superAdmin />
    </>
  );
};
