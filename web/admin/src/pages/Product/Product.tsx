import { getProducts } from "@/slices/product/productSlice";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { useDispatch, useSelector } from "react-redux";
import { ProductClient } from "./components/client";
import { myApi } from "@/helpers/api";

export default function ProductsPage() {
  const dispatch = useDispatch();
  const { products } = useSelector((state: RootState) => state.product);

  const { isLoading } = useQuery({
    queryKey: ["products"],
    queryFn: async () => {
      const { data } = await myApi.get("/product");
      dispatch(getProducts(data.data));
      return data.data;
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <ProductClient data={products} />
      </div>
    </div>
  );
}
