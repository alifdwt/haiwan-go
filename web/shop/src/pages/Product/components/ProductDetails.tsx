import StarRating from "@/components/star";
import { Product } from "@/interface/Product";
import { ToRupiah } from "@/lib/currency";

export default function ProductDetails({ product }: { product: Product }) {
  return (
    <>
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
    </>
  );
}
