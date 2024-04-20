// {
//     "id": 7,
//     "name": "Udyliv Mmmlma",
//     "price": "962",
//     "image": "https://picsum.photos/id/57/720",
//     "quantity": 3,
//     "weight": 832,
//     "user_id": 1,
//     "product_id": 4,
//     "created_at": "2024-04-20T17:09:09.593796+07:00",
//     "updated_at": "2024-04-20T17:09:09.593796+07:00"
// }

export interface Cart {
  id: number;
  name: string;
  price: string;
  image: string;
  quantity: number;
  weight: number;
  user_id: number;
  product_id: number;
  created_at: string;
  updated_at: string;
}
