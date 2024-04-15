// {
//     "id": 1,
//     "name": "Beauty Cat Food",
//     "category_id": 3,
//     "description": "Diformulasikan untuk adult & kitten, Menggunakan bahan baku natural yang nutrisi nya diformulasikan secara seimbang.",
//     "price": 27000,
//     "count_in_stock": 25,
//     "brand": "Beauty",
//     "weight": 1000,
//     "rating": 5,
//     "slug": "beauty-cat-food",
//     "image_path": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1713193373/HaiwanGo/beauty-cat-food-1.webp",
//     "created_at": "2024-04-15 22:02:53.670873 +0700 WIB",
//     "updated_at": "2024-04-15 22:02:53.670873 +0700 WIB",
//     "category": {
//       "id": 3,
//       "name": "Cat",
//       "description": "Lorem ipsum dolor sit amet consectetur adipiscing elit semper dignissim accumsan nunc duis fames id volutpat in eros luctus phasellus molestie aptent nec lacinia aliquet cras justo parturient ad.",
//       "slug": "cat",
//       "image_path": "https://res.cloudinary.com/dxirtmo5t/image/upload/v1713189961/HaiwanGo/cat.png",
//       "created_at": "2024-04-15 21:06:02.109856 +0700 WIB",
//       "updated_at": "2024-04-15 21:06:02.109856 +0700 WIB"
//     }
//   }

import { Category } from "./Category";

export interface Product {
  id: number;
  name: string;
  category_id: number;
  description: string;
  price: number;
  count_in_stock: number;
  brand: string;
  weight: number;
  rating: number;
  slug: string;
  image_path: string;
  created_at: string;
  updated_at: string;
  category?: Category;
}
