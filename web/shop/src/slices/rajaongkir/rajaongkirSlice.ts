import { myApi } from "@/helpers/api";
import { RootState } from "@/store";
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { useSelector } from "react-redux";

const initialState = {
  provinces: [],
  cities: [],
  shippingCosts: [],
  loading: false,
  error: null,
};

export const fetchProvinces = createAsyncThunk(
  "rajaOngkir/fetchProvinces",
  async (_, { rejectWithValue }) => {
    try {
      const { access_token } = useSelector((state: RootState) => state.user);
      if (!access_token) {
        console.error("Access token not found");
        return rejectWithValue("Access token not found");
      }

      const response = await myApi.get("/rajaongkir/provinsi", {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });

      return response.data.rajaongkir.results;
    } catch (error) {
      // @ts-expect-error next-line
      return rejectWithValue(error.response.data);
    }
  }
);

export const fetchCities = createAsyncThunk(
  "rajaOngkir/fetchCities",
  async (provinceId: number, { rejectWithValue }) => {
    try {
      const { access_token } = useSelector((state: RootState) => state.user);
      if (!access_token) {
        console.error("Access token not found");
        return rejectWithValue("Access token not found");
      }

      const response = await myApi.get(`/rajaongkir/city/${provinceId}`, {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      });

      return response.data.rajaongkir.results;
    } catch (error) {
      // @ts-expect-error next-line
      return rejectWithValue(error.response.data);
    }
  }
);

export const calculateShippingCost = createAsyncThunk(
  "rajaOngkir/calculateShippingCost",
  async (
    {
      asal,
      tujuan,
      berat,
      kurir,
    }: { asal: string; tujuan: string; berat: number; kurir: string },
    { rejectWithValue }
  ) => {
    try {
      const { access_token } = useSelector((state: RootState) => state.user);

      if (!access_token) {
        console.error("Access token not found");
        return rejectWithValue("Access token not found");
      }

      const response = await myApi.post(
        "/rajaongkir/cost",
        {
          asal,
          tujuan,
          berat,
          kurir,
        },
        {
          headers: {
            Authorization: `Bearer ${access_token}`,
          },
        }
      );

      return response.data.rajaongkir.results;
    } catch (error) {
      // @ts-expect-error next-line
      return rejectWithValue(error.response.data);
    }
  }
);

const rajaOngkirSlice = createSlice({
  name: "rajaOngkir",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchProvinces.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchProvinces.fulfilled, (state, action) => {
        state.loading = false;
        state.provinces = action.payload;
      })
      .addCase(fetchProvinces.rejected, (state, action) => {
        state.loading = false;
        // @ts-expect-error next-line
        state.error = action.error.message;
      })
      .addCase(fetchCities.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchCities.fulfilled, (state, action) => {
        state.loading = false;
        state.cities = action.payload;
      })
      .addCase(fetchCities.rejected, (state, action) => {
        state.loading = false;
        // @ts-expect-error next-line
        state.error = action.error.message;
      })
      .addCase(calculateShippingCost.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(calculateShippingCost.fulfilled, (state, action) => {
        state.loading = false;
        state.shippingCosts = action.payload;
      })
      .addCase(calculateShippingCost.rejected, (state, action) => {
        state.loading = false;
        // @ts-expect-error next-line
        state.error = action.error.message;
      });
  },
});

export default rajaOngkirSlice.reducer;
