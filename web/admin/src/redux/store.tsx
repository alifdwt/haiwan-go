import { configureStore } from "@reduxjs/toolkit";
import { deleteUserSlice, getAllUsersSlice, updateSlice } from "./user.slice";
import { loginSlice, registerSlice } from "./auth.slice";

const store = configureStore({
  reducer: {
    updateUserReducer: updateSlice.reducer,
    getAllusersReducer: getAllUsersSlice.reducer,
    deleteUserReducer: deleteUserSlice.reducer,
    loginReducer: loginSlice.reducer,
    registerReducer: registerSlice.reducer,
  },
});

export default store;
