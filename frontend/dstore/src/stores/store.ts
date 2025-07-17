import { configureStore } from "@reduxjs/toolkit";
import cartReducer from "../features/order/cartSlice";

export const store = configureStore({
  reducer: {
    cart: cartReducer,
  },
});


export type AppState = typeof store;
export type RootState = ReturnType<AppState["getState"]>;
export default store;