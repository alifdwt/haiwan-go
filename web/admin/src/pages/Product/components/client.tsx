import { Heading } from "@/components/ui/heading";
import { ProductColumns, columns } from "./columns";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { DataTable } from "@/components/ui/data-table";
import { ApiList } from "@/components/ui/api-list";

interface ProductClientProps {
  data: ProductColumns[];
}

export const ProductClient: React.FC<ProductClientProps> = ({ data }) => {
  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Produk (${data.length})`}
          description="Lihat dan atur produkmu"
        />
        <a href="/products/create">
          <Button variant={"secondary"}>+ Tambah Produk</Button>
        </a>
      </div>
      <Separator />
      <DataTable searchKey="name" columns={columns} data={data} />
      <Heading title="API" description="Dokumentasi API untuk produk" />
      <Separator />
      <ApiList entityName="product" entityIdName="productId" superAdmin />
    </>
  );
};
