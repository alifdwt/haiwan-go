import { Toaster } from "@/components/ui/toaster";
import Layout from "@/layouts";
import LandingLayout from "@/layouts/Landing/Landing";
import Home from "@/pages/Home";
import CreateSliderPage from "@/pages/Slider/Create/Create";
import SliderPage from "@/pages/Slider/Slider";
import LoginPage from "@/pages/auth/Login";
import RegisterPage from "@/pages/auth/Register";
import { RootState } from "@/store";
import { useSelector } from "react-redux";
import {
  BrowserRouter,
  Navigate,
  Route,
  Routes as RouterRoutes,
} from "react-router-dom";

export default function Routes() {
  const { access_token } = useSelector((state: RootState) => state.user);
  return (
    <BrowserRouter>
      <RouterRoutes>
        <Route element={access_token ? <Layout /> : <Navigate to="/" />}>
          <Route path="/dashboard" element={<div>Dashboard</div>} />
          <Route path="/sliders" element={<SliderPage />} />
          <Route path="/sliders/create" element={<CreateSliderPage />} />
        </Route>
        <Route
          element={
            access_token ? <Navigate to="/dashboard" /> : <LandingLayout />
          }
        >
          <Route path="/" element={<Home />} />
          <Route
            path="/login"
            element={access_token ? <Navigate to="/" /> : <LoginPage />}
          />
          <Route
            path="/register"
            element={access_token ? <Navigate to="/" /> : <RegisterPage />}
          />
        </Route>
      </RouterRoutes>
      <Toaster />
    </BrowserRouter>
  );
}
