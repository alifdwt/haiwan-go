import Layout from "@/layouts";
import Home from "@/pages/Home/Home";
import { BrowserRouter, Route, Routes as RouterRoutes } from "react-router-dom";

export default function Routes() {
  return (
    <BrowserRouter>
      <RouterRoutes>
        <Route element={<Layout />}>
          <Route path="/" element={<Home />} />
        </Route>
      </RouterRoutes>
    </BrowserRouter>
  );
}
