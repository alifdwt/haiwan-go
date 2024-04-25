import { Product } from "./Product";

export interface Category {
  id: number;
  name: string;
  description: string;
  slug: string;
  image_path: string;
  created_at: string;
  updated_at: string;
}

export interface CategoryWithRelation {
  id: number;
  name: string;
  description: string;
  slug: string;
  image_path: string;
  created_at: string;
  updated_at: string;
  products: Product[];
}
