import { Button } from "@/components/ui/button";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { useToast } from "@/components/ui/use-toast";
import { logout } from "@/slices/user/userSlice";
import { RootState } from "@/store";
import { useDispatch, useSelector } from "react-redux";

export default function LogOutButton() {
  const { user } = useSelector((state: RootState) => state.user);
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
    <Popover>
      <PopoverTrigger asChild>
        <Button className="ml-auto" variant={"destructive"}>
          {user?.firstname}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-56">
        <Button
          variant={"destructive"}
          className="w-full"
          onClick={handleLogout}
        >
          Logout
        </Button>
      </PopoverContent>
    </Popover>
  );
}
