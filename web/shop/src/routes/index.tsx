import { Toaster } from "@/components/ui/toaster";
import Layout from "@/layouts";
import CartPage from "@/pages/Cart/Cart";
import CategoryPage from "@/pages/Category/Category";
import Home from "@/pages/Home/Home";
import LoginPage from "@/pages/Login/Login";
import ProductPage from "@/pages/Product/Product";
import ShippingPage from "@/pages/Shipping/Shipping";
import AuthProvider from "@/slices/AuthProvider";
import { BrowserRouter, Route, Routes as RouterRoutes } from "react-router-dom";

export default function Routes() {
  return (
    <BrowserRouter>
      <RouterRoutes>
        <Route element={<Layout />}>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/" element={<Home />} />
          <Route path="/products/:productSlug" element={<ProductPage />} />
          <Route path="/categories/:categorySlug" element={<CategoryPage />} />
          <Route
            path="/cart"
            element={
              <AuthProvider>
                <CartPage />
              </AuthProvider>
            }
          />
          <Route
            path="/cart/shipment"
            element={
              <AuthProvider>
                <ShippingPage />
              </AuthProvider>
            }
          />
        </Route>
      </RouterRoutes>
      <Toaster />
    </BrowserRouter>
  );
}
