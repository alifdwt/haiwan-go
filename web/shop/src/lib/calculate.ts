import { Cart } from "@/interface/Cart";

export const calculateSubtotal = (cartItems: Cart[]) => {
  return cartItems.reduce((total, item) => {
    return total + item.quantity * parseInt(item.price);
  }, 0);
};

export const calculateTotalProducts = (cartItems: Cart[]) => {
  return cartItems.reduce((total, item) => {
    return total + item.quantity;
  }, 0);
};

export const calculateWeight = (cartItems: Cart[]) => {
  return cartItems.reduce((totalWeight, item) => {
    return totalWeight + item.quantity * item.weight;
  }, 0);
};
