import { Toaster } from "@/components/ui/toaster";
import Layout from "@/layouts";
import CartPage from "@/pages/Cart/Cart";
import Home from "@/pages/Home/Home";
import ProductPage from "@/pages/Product/Product";
import ShippingPage from "@/pages/Shipping/Shipping";
import { RootState } from "@/store";
import { useSelector } from "react-redux";
import { BrowserRouter, Route, Routes as RouterRoutes } from "react-router-dom";

export default function Routes() {
  const { access_token } = useSelector((state: RootState) => state.user);
  return (
    <BrowserRouter>
      <RouterRoutes>
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
          <Route path="/products/:productSlug" element={<ProductPage />} />
          <Route
            path="/cart"
            element={access_token ? <CartPage /> : <Home />}
          />
          <Route
            path="/cart/shipment"
            element={access_token ? <ShippingPage /> : <Home />}
          />
        </Route>
      </RouterRoutes>
      <Toaster />
    </BrowserRouter>
  );
}
