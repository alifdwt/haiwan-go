import { Outlet } from "react-router-dom";
import Navbar from "./Navbar/Navbar";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { myApi } from "@/helpers/api";
import { logout, updateTokenAsync } from "@/slices/user/userSlice";
import { useToast } from "@/components/ui/use-toast";

export default function Layout() {
  const { toast } = useToast();
  const refresh_token = useSelector(
    (state: RootState) => state.user.refresh_token
  );
  const dispatch = useDispatch();

  const { isError } = useQuery({
    queryKey: ["refresh"],
    queryFn: async () => {
      if (refresh_token) {
        const response = await myApi.post("/auth/refresh-token", {
          refresh_token: refresh_token,
        });
        dispatch(updateTokenAsync(response.data.data));
        return response;
      }
    },
  });

  if (isError) {
    dispatch(logout());
    toast({
      title: "Sesi anda telah habis",
      description: "Silahkan login kembali",
      variant: "destructive",
    });
  }

  return (
    <div className="min-h-screen">
      <Navbar />
      <main className="p-4">
        <Outlet />
      </main>
    </div>
  );
}
