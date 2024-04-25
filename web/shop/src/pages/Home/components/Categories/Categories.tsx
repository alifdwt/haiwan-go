import { myApi } from "@/helpers/api";
import { getCategories } from "@/slices/category/categorySlice";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { useDispatch, useSelector } from "react-redux";

export default function Categories() {
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

  // console.log(data);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="grid grid-cols-2 md:grid-cols-8 gap-4">
      {categories.map((category) => (
        <a
          href={`/categories/${category.slug}`}
          key={category.id}
          className="flex flex-col items-center gap-2 border border-gray-300 rounded-lg p-4 w-36"
        >
          <img
            className="object-cover"
            src={category.image_path}
            alt={category.name}
          />
          <p className="truncate">{category.name}</p>
        </a>
      ))}
    </div>
  );
}
