import { myApi } from "@/helpers/api";
import { Product } from "@/interface/Product";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import ProductDetails from "./components/ProductDetails";
import ProductTransaction from "./components/ProductTransaction";

export default function ProductPage() {
  const { productSlug } = useParams();

  const { data, isLoading } = useQuery({
    queryKey: ["product", productSlug],
    queryFn: async () => {
      const { data } = await myApi.get(`/product/slug/${productSlug}`);
      return data.data;
    },
  });

  const product: Product = data;

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="max-w-screen-xl mx-auto mb-5 p-2">
      <div className="grid grid-cols-1 md:grid-cols-[1fr_.75fr]">
        <div className="col-span-1 mx-auto md:w-[500px]">
          <img
            src={product.image_path}
            alt={product.name}
            className="max-h-[500px] object-cover rounded-lg"
          />
        </div>
        <div className="col-span-1 flex flex-col gap-4">
          <ProductDetails product={product} />
          <ProductTransaction product={product} />
          <table className="table-fixed">
            <thead>
              <tr>
                <th className="text-left">Category</th>
                <th className="text-left">Brand</th>
                <th className="text-left">Weight</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>{product.category?.name}</td>
                <td>{product.brand}</td>
                <td>{product.weight}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
