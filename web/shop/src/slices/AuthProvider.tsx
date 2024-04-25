import { RootState } from "@/store";
import { useQuery } from "@tanstack/react-query";
import { useDispatch, useSelector } from "react-redux";
import { useNavigate } from "react-router-dom";
import { logout, updateTokenAsync } from "./user/userSlice";
import { myApi } from "@/helpers/api";
import { useToast } from "@/components/ui/use-toast";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const { toast } = useToast();
  const { refresh_token } = useSelector((state: RootState) => state.user);
  const dispatch = useDispatch();
  const navigate = useNavigate();

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
      title: "Silakan login",
      description: "Silakan login untuk mengakses halaman ini",
      variant: "destructive",
    });
    navigate("/login");
  }

  return <>{children}</>;
}
