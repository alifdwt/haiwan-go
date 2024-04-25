import { Category } from "@/interface/Category";
import { createSlice } from "@reduxjs/toolkit";

interface CategorySliceState {
  categories: Category[];
}

const initialState: CategorySliceState = {
  categories: [],
};

export const categorySlice = createSlice({
  name: "category",
  initialState,
  reducers: {
    getCategories: (state, action) => {
      state.categories = action.payload;
    },
  },
});

export const { getCategories } = categorySlice.actions;

export default categorySlice.reducer;
