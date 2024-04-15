import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import { CategoryForm } from "./components/category-form";

export default function CategoryActionPage() {
  const { categoryId } = useParams();
  const { access_token } = useSelector((state: RootState) => state.user);

  const { data, isLoading } = useQuery({
    queryKey: ["category", categoryId],
    queryFn: async () => {
      const { data } = await myApi.get(`/category/${categoryId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });
      return data.data;
    },
  });

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        {categoryId === "create" ? (
          <CategoryForm initialData={null} />
        ) : isLoading ? (
          <div>Loading...</div>
        ) : (
          <CategoryForm initialData={data} />
        )}
      </div>
    </div>
  );
}
