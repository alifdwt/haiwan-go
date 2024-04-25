import ProductCard from "@/components/product-card";
import { myApi } from "@/helpers/api";
import { CategoryWithRelation } from "@/interface/Category";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";

export default function CategoryPage() {
  const { categorySlug } = useParams();

  const { data, isLoading } = useQuery({
    queryKey: ["category", categorySlug],
    queryFn: async () => {
      const { data } = await myApi.get(`/category/slug/${categorySlug}`);
      return data.data;
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  const category: CategoryWithRelation = data;
  console.log(category);

  return (
    <>
      <div className="relative w-[80vw]">
        <img
          src="https://picsum.photos/id/4/32/32"
          alt={category.name}
          className="w-full h-72 object-cover rounded-lg"
        />
        <p className="text-3xl font-bold absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
          {category.name}
        </p>
        <p className="text-sm absolute bottom-1/4 left-1/2 -translate-x-1/2 text-center">
          {category.description}
        </p>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-4 mt-5">
        <div className="col-span-1">
          <h3 className="text-3xl font-bold">Products</h3>
        </div>
        <div className="col-span-3">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            {category.products &&
              category.products.map((product) => (
                <ProductCard key={product.id} product={product} />
              ))}
          </div>
        </div>
      </div>
    </>
  );
}
