import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { myApi } from "../helpers/api";

interface UpdateUserPayload {
  userId: string;
  updateUser: {
    firstname: string;
    lastname: string;
    email: string;
    phone: string;
  };
}

export const updateUser = createAsyncThunk(
  "user/updateUser",
  async ({ userId, updateUser }: UpdateUserPayload, { rejectWithValue }) => {
    try {
      const response = await myApi.put(`/users/${userId}`, updateUser);
      return response.data;
    } catch (error) {
      return rejectWithValue(error);
    }
  }
);

export const getAllUsers = createAsyncThunk(
  "user/getAllUsers",
  async (_, { rejectWithValue }) => {
    try {
      const response = await myApi.get("/users");
      return response.data;
    } catch (error) {
      return rejectWithValue(error);
    }
  }
);

export const deleteUser = createAsyncThunk(
  "user/deleteUser",
  async (_, { rejectWithValue }) => {
    try {
      const response = await myApi.get("/users");
      return response.data;
    } catch (error) {
      return rejectWithValue(error);
    }
  }
);

export const updateSlice = createSlice({
  name: "update",
  initialState: {
    loading: false,
    success: false,
    error: "",
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(updateUser.pending, (state) => {
        state.loading = true;
        state.success = false;
        state.error = "";
      })
      .addCase(updateUser.fulfilled, (state) => {
        state.loading = false;
        state.success = true;
        // state.error = "";
      })
      .addCase(updateUser.rejected, (state, action) => {
        state.loading = false;
        // state.success = false;
        state.error = action.error.message as string;
      });
  },
});

export const getAllUsersSlice = createSlice({
  name: "getAllUsers",
  initialState: {
    loading: false,
    users: [],
    error: "",
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getAllUsers.pending, (state) => {
        state.loading = true;
        state.users = [];
        state.error = "";
      })
      .addCase(getAllUsers.fulfilled, (state, action) => {
        state.loading = false;
        state.users = action.payload;
        // state.error = "";
      })
      .addCase(getAllUsers.rejected, (state, action) => {
        state.loading = false;
        // state.users = [];
        state.error = action.error.message as string;
      });
  },
});

export const deleteUserSlice = createSlice({
  name: "deleteUser",
  initialState: {
    loading: false,
    success: false,
    error: "",
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(deleteUser.pending, (state) => {
        state.loading = true;
        state.success = false;
        state.error = "";
      })
      .addCase(deleteUser.fulfilled, (state) => {
        state.loading = false;
        state.success = true;
        // state.error = "";
      })
      .addCase(deleteUser.rejected, (state, action) => {
        state.loading = false;
        // state.success = false;
        state.error = action.error.message as string;
      });
  },
});

export default {
  updateSlice: updateSlice.reducer,
  getAllUsersSlice: getAllUsersSlice.reducer,
  deleteUserSlice: deleteUserSlice.reducer,
};
