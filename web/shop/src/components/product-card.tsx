import { Product } from "@/interface/Product";
import { ToRupiah } from "@/lib/currency";
import StarRating from "./star";

export default function ProductCard({ product }: { product: Product }) {
  return (
    <a href={`/products/${product.slug}`} className="border rounded-lg">
      <div>
        <img
          src={product.image_path}
          alt={product.name}
          className="w-full h-full object-cover rounded-t-lg"
        />
        <div className="p-2">
          <p className="font-bold">{product.name}</p>
          <p>{ToRupiah(product.price)}</p>
          <StarRating rating={product.rating} />
          <p className="text-xs text-muted-foreground">
            Tersisa {product.count_in_stock} lagi
          </p>
        </div>
      </div>
    </a>
  );
}
