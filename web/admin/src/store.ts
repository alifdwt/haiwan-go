import { combineReducers, configureStore } from "@reduxjs/toolkit";
import { persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";

import UserReducer from "@/slices/user/userSlice";
import SliderReducer from "@/slices/slider/sliderSlice";
import CategoryReducer from "@/slices/category/categorySlice";

const persistConfig = {
  key: "root",
  storage,
};

const rootReducer = combineReducers({
  user: UserReducer,
  slider: SliderReducer,
  category: CategoryReducer,
});

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
});

export const persistor = persistStore(store);
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
