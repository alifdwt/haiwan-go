import React from "react";
import { SliderColumns, columns } from "./columns";
import { Heading } from "@/components/ui/heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { DataTable } from "@/components/ui/data-table";
import { ApiList } from "@/components/ui/api-list";

interface SliderClientProps {
  data: SliderColumns[];
}

export const SliderClient: React.FC<SliderClientProps> = ({ data }) => {
  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Slider (${data.length})`}
          description="Lihat dan atur slidermu"
        />
        <a href="/sliders/create">
          <Button variant={"secondary"}>+ Tambah Slider</Button>
        </a>
      </div>
      <Separator />
      <DataTable searchKey="name" columns={columns} data={data} />
      <Heading title="API" description="Dokumentasi API untuk slider" />
      <Separator />
      <ApiList entityName="slider" entityIdName="sliderId" />
    </>
  );
};
