import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { login, logout } from "@/slices/user/userSlice";
import { RootState } from "@/store";
import { useForm } from "react-hook-form";
import { useDispatch, useSelector } from "react-redux";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { myApi } from "@/helpers/api";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const loginFormSchema = z.object({
  email: z.string().min(1, { message: "Email is required" }),
  password: z.string().min(1, { message: "Password is required" }),
});

type LoginFormValues = z.infer<typeof loginFormSchema>;

export default function UserNav() {
  const { user } = useSelector((state: RootState) => state.user);
  const dispatch = useDispatch();

  const loginForm = useForm<LoginFormValues>({
    defaultValues: {
      email: "",
      password: "",
    },
    resolver: zodResolver(loginFormSchema),
  });

  const mutation = useMutation({
    mutationFn: (data: LoginFormValues) => {
      return myApi.post("/auth/login", data);
    },
  });

  const handleLogin = (data: LoginFormValues) => {
    mutation.mutate(data, {
      onSuccess: (data) => {
        dispatch(login(data.data.data));
        window.location.reload();
      },
      onError: (error) => {
        console.log(error);
      },
    });
  };

  const handleLogout = () => {
    dispatch(logout());
  };

  if (user) {
    return (
      <Popover>
        <PopoverTrigger asChild>
          <Button className="ml-auto" variant={"secondary"}>
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

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="ml-auto" variant={"secondary"}>
          Masuk
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Masuk</DialogTitle>
          <DialogDescription>
            Masuk ke akun anda untuk melanjutkan transaksi
          </DialogDescription>
        </DialogHeader>
        <Form {...loginForm}>
          <form onSubmit={loginForm.handleSubmit(handleLogin)}>
            <FormField
              control={loginForm.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      disabled={mutation.isPending}
                      {...field}
                      placeholder="Email"
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={loginForm.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input
                      disabled={mutation.isPending}
                      {...field}
                      placeholder="Password"
                      type="password"
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button
              type="submit"
              disabled={mutation.isPending}
              className="w-full mt-4"
            >
              Masuk
            </Button>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  );
}
