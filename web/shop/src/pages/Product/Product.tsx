import StarRating from "@/components/star";
import { Button } from "@/components/ui/button";
import { myApi } from "@/helpers/api";
import { Product } from "@/interface/Product";
import { ToRupiah } from "@/lib/currency";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { useSelector } from "react-redux";
import { useParams } from "react-router-dom";

export default function ProductPage() {
  const { productSlug } = useParams();
  const { access_token } = useSelector((state: RootState) => state.user);

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
            className="w-full h-full object-cover rounded-lg"
          />
        </div>
        <div className="col-span-1 flex flex-col gap-4">
          <h1 className="text-3xl font-extrabold">{product.name}</h1>
          <div className="flex gap-2">
            <StarRating rating={product.rating} />
            <p className="text-sm font-medium">
              {product.rating} <span className="text-primary">/ 5</span>
            </p>
            <p className="text-sm text-secondaryDark">
              {product.count_in_stock} tersisa
            </p>
          </div>
          <div className="flex gap-2 items-center">
            <p className="text-xl font-bold text-primary">
              {ToRupiah(product.price)}
            </p>
            <p className="text-sm text-secondaryDark line-through">
              {ToRupiah(product.price * 1.5)}
            </p>
          </div>
          <p className="text-sm">{product.description}</p>
          <div className="flex gap-2">
            {access_token && (
              <Button variant={"secondary"} className="w-full">
                Masukkan ke Keranjang
              </Button>
            )}
            <Button className="w-full">Beli</Button>
          </div>
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
