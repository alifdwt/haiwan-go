import { decodeToken } from "@/helpers/decode";
import User from "@/interface/User";
import { createSlice } from "@reduxjs/toolkit";

interface UserSliceState {
  user?: User;
  access_token?: string;
  refresh_token?: string;
}

const initialState: UserSliceState = {
  user: undefined,
  access_token: undefined,
  refresh_token: undefined,
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    login: (state, action) => {
      const decodedUser = decodeToken(action.payload.access_token);
      state.user = decodedUser !== null ? decodedUser : undefined;
      state.access_token = action.payload.access_token;
      state.refresh_token = action.payload.refresh_token;
    },
    setUser: (state, action) => {
      const decodedUser = decodeToken(action.payload.access_token);
      state.user = decodedUser !== null ? decodedUser : undefined;
    },
    updateTokenAsync: (state, action) => {
      state.access_token = action.payload.access_token;
      state.refresh_token = action.payload.refresh_token;
    },
    logout: (state) => {
      state.user = undefined;
      state.access_token = undefined;
      state.refresh_token = undefined;
    },
  },
});

export const { login, setUser, updateTokenAsync, logout } = userSlice.actions;

export default userSlice.reducer;
