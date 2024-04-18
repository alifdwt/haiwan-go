import { useForm } from "react-hook-form";
import * as z from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { myApi } from "@/helpers/api";
import { useNavigate } from "react-router-dom";
import { useToast } from "@/components/ui/use-toast";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useMutation } from "@tanstack/react-query";

const formSchema = z.object({
  name: z.string().min(1, { message: "Name is required" }),
  email: z.string().min(1, { message: "Email is required" }),
  password: z.string().min(1, { message: "Password is required" }),
  confirm_password: z
    .string()
    .min(1, { message: "Confirm password is required" }),
});

type RegisterFormValues = z.infer<typeof formSchema>;

interface RegisterFormProps {
  initialData: null;
}

export const RegisterForm: React.FC<RegisterFormProps> = ({ initialData }) => {
  const navigate = useNavigate();
  const { toast } = useToast();

  const form = useForm<RegisterFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: initialData || {},
  });

  const mutation = useMutation({
    mutationFn: (data: RegisterFormValues) => {
      return myApi.post("/auth/register", data);
    },
  });

  const onSubmit = (data: RegisterFormValues) => {
    mutation.mutate(data, {
      onSuccess: () => {
        toast({
          title: "Register berhasil",
          description: "Selamat datang kembali",
        });
        navigate("/login");
      },
      onError: (error) => {
        console.log(error);
        toast({
          title: "Register gagal",
          // @ts-expect-error next-line
          description: `${error.response.data.message}`,
        });
      },
    });
  };

  return (
    <>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-white">Name</FormLabel>
                <FormControl>
                  <Input
                    disabled={mutation.isPending}
                    placeholder="Isi dengan nama Anda"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-white">Email</FormLabel>
                <FormControl>
                  <Input
                    disabled={mutation.isPending}
                    placeholder="Isi dengan email Anda"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="grid grid-cols-2 gap-4">
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-white">Password</FormLabel>
                  <FormControl>
                    <Input
                      disabled={mutation.isPending}
                      type="password"
                      placeholder="Isi dengan password Anda"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="confirm_password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-white">Confirm Password</FormLabel>
                  <FormControl>
                    <Input
                      disabled={mutation.isPending}
                      type="password"
                      placeholder="Isi dengan password Anda"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <Button
            disabled={mutation.isPending}
            type="submit"
            className="w-full"
            variant={"secondary"}
          >
            <span>Daftar</span>
            {/* <ArrowRightFromLine /> */}
          </Button>
        </form>
      </Form>
    </>
  );
};
