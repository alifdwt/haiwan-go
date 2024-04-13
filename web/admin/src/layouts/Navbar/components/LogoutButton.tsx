import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { logout } from "@/slices/user/userSlice";
import { useDispatch } from "react-redux";

export default function LogOutButton() {
  const dispatch = useDispatch();
  const { toast } = useToast();

  const handleLogout = () => {
    dispatch(logout());
    toast({
      title: "Logout Berhasil",
      description: "Anda Berhasil Logout",
      variant: "destructive",
    });
  };
  return (
    <Button className="ml-auto" variant={"destructive"} onClick={handleLogout}>
      Keluar
    </Button>
  );
}
