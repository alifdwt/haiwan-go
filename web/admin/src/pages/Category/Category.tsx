import { useDispatch, useSelector } from "react-redux";
import { CategoryClient } from "./components/client";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { myApi } from "@/helpers/api";
import { getCategories } from "@/slices/category/categorySlice";

export default function CategoryPage() {
  const dispatch = useDispatch();
  const { categories } = useSelector((state: RootState) => state.category);

  const { isLoading } = useQuery({
    queryKey: ["categories"],
    queryFn: async () => {
      const { data } = await myApi.get("/category");
      dispatch(getCategories(data.data));
      return data.data;
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <CategoryClient data={categories} />
      </div>
    </div>
  );
}
