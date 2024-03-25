import jwt from "jsonwebtoken";
import { createSlice, createAsyncThunk, Dispatch } from "@reduxjs/toolkit";
import { myApi } from "../helpers/api";

type MyJwtPayload = {
  firstname: string;
  lastname: string;
  email: string;
  phone: string;
};

let accessToken: string | null = null;
let refreshToken: string | null = null;

if (typeof window !== "undefined") {
  accessToken = localStorage.getItem("accessToken");
  refreshToken = localStorage.getItem("refreshToken");
}

const initialUser = accessToken
  ? (jwt.decode(accessToken) as MyJwtPayload)
  : null;

export const registerAction = createAsyncThunk(
  "auth/register",
  async (registerData, { rejectWithValue }) => {
    try {
      const response = await myApi.post("/auth/register", registerData);
      return response.data;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      } else {
        return rejectWithValue("Unexpected error");
      }
    }
  }
);

export const loginAction = createAsyncThunk(
  "auth/login",
  async (loginData, { rejectWithValue }) => {
    try {
      const response = await myApi.post("/auth/login", loginData);

      localStorage.setItem("accessToken", response.data.data.accessToken);

      localStorage.setItem("refreshToken", response.data.data.refreshToken);

      return response.data.data;
    } catch (error) {
      if (error instanceof Error) {
        return rejectWithValue(error.message);
      } else {
        return rejectWithValue("Unexpected error");
      }
    }
  }
);

export const logoutUserAction =
  (navigate: (path: string) => void) => (dispatch: Dispatch) => {
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
    dispatch(clearAuthData());
    navigate("/login");
  };

// export const updateTokenAsync = createAsyncThunk(
//   "auth/updateToken",
//   async (_, { getState, dispatch }) => {
//     try {
//       const { refreshToken } = getState().authReducer;

//       if (!refreshToken) {
//         return;
//       }

//       console.log(refreshToken);

//       const response = await myApi.post("/auth/refresh", {
//         refresh: refreshToken,
//       });

//       const data = response.data;

//       localStorage.setItem("accessToken", data);
//     } catch (error) {
//       console.log(error);
//     }
//   }
// );

export const registerSlice = createSlice({
  name: "register",
  initialState: {
    loading: false,
    success: false,
    error: "",
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(registerAction.pending, (state) => {
        state.loading = true;
        state.error = "";
      })
      .addCase(registerAction.fulfilled, (state) => {
        state.loading = false;
        state.success = true;
      })
      .addCase(registerAction.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });
  },
});

export const loginSlice = createSlice({
  name: "auth",
  initialState: {
    user: initialUser,
    accessToken: accessToken,
    refreshToken: refreshToken,
    loading: false,
    error: "",
    authTokens: null,
  },
  reducers: {
    setUser: (state, action) => {
      state.user = action.payload;
    },
    setAuthTokens: (state, action) => {
      state.authTokens = action.payload;
    },
    setLoading: (state, action) => {
      state.loading = action.payload;
    },
    clearAuthData: (state) => {
      state.user = null;
      state.authTokens = null;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(registerAction.pending, (state) => {
        state.loading = true;
        state.error = "";
      })
      .addCase(loginAction.fulfilled, (state, action) => {
        state.loading = false;
        state.accessToken = action.payload.data.accessToken;
        state.refreshToken = action.payload.data.refreshToken;

        state.user = action.payload.data.user;
      })
      .addCase(loginAction.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message as string;
      });
  },
});

export const { setUser, setAuthTokens, setLoading, clearAuthData } =
  loginSlice.actions;

export default {
  setUser: setUser,
  setAuthTokens: setAuthTokens,
  setLoading: setLoading,
  clearAuthData: clearAuthData,
  loginSlice: loginSlice.reducer,
  registerSlice: registerSlice.reducer,
};
