import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useToast } from "@/components/ui/use-toast";
import { myApi } from "@/helpers/api";
import { login } from "@/slices/user/userSlice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useMutation } from "@tanstack/react-query";
import { useForm } from "react-hook-form";
import { useDispatch } from "react-redux";
import { useNavigate } from "react-router-dom";
import * as z from "zod";

const formSchema = z.object({
  email: z.string().min(1, { message: "Email is required" }),
  password: z.string().min(1, { message: "Password is required" }),
});

type LoginFormValues = z.infer<typeof formSchema>;

interface LoginFormProps {
  initialData: null;
}

export const LoginForm: React.FC<LoginFormProps> = ({ initialData }) => {
  const { toast } = useToast();
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const form = useForm<LoginFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: initialData || {},
  });

  const mutation = useMutation({
    mutationFn: (data: LoginFormValues) => {
      return myApi.post("/auth/login", data);
    },
  });

  const onSubmit = (data: LoginFormValues) => {
    mutation.mutate(data, {
      onSuccess: (data) => {
        toast({
          title: "Login berhasil",
          description: "Selamat datang kembali",
        });
        dispatch(login(data.data.data));
        navigate("/");
      },
      onError: (error) => {
        toast({
          title: "Login gagal",
          description: `${error.message}`,
        });
      },
    });
  };

  return (
    <>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    disabled={mutation.isPending}
                    placeholder="Email"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input
                    disabled={mutation.isPending}
                    type="password"
                    placeholder="Password"
                    {...field}
                  />
                </FormControl>
              </FormItem>
            )}
          />
          <Button
            type="submit"
            className="w-full mt-4 hover:bg-black"
            disabled={mutation.isPending}
          >
            Masuk
          </Button>
        </form>
      </Form>
    </>
  );
};
