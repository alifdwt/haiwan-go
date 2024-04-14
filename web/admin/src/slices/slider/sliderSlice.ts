import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  sliders: [],
};

export const sliderSlice = createSlice({
  name: "slider",
  initialState,
  reducers: {
    getSliders: (state, action) => {
      state.sliders = action.payload;
    },
  },
});

export const { getSliders } = sliderSlice.actions;

export default sliderSlice.reducer;
