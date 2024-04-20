import { createSlice } from "@reduxjs/toolkit";
import { Cart } from "@/interface/Cart";

interface CartSliceState {
  carts: Cart[];
}

const initialState: CartSliceState = {
  carts: [],
};

export const cartSlice = createSlice({
  name: "cart",
  initialState,
  reducers: {
    getCarts: (state, action) => {
      state.carts = action.payload;
    },
  },
});

export const { getCarts } = cartSlice.actions;

export default cartSlice.reducer;
