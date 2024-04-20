import { Cart } from "@/interface/Cart";

function getTotalPrice(cart: Cart[]): number {
  let totalPrice: number = 0;

  cart.forEach((item) => {
    const itemPrice: number = parseInt(item.price) * item.quantity;
    totalPrice += itemPrice;
  });

  return totalPrice;
}

function getTotalQuantity(cart: Cart[]): number {
  let totalQuantity: number = 0;

  cart.forEach((item) => {
    totalQuantity += item.quantity;
  });

  return totalQuantity;
}

export { getTotalPrice, getTotalQuantity };
