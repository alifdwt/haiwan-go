import React, { useState } from "react";
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

const formSchema = z.object({
  name: z.string().min(1, { message: "Name is required" }),
  email: z.string().min(1, { message: "Email is required" }),
  password: z.string().min(1, { message: "Password is required" }),
  confirmPassword: z
    .string()
    .min(1, { message: "Confirm password is required" }),
});

type RegisterFormValues = z.infer<typeof formSchema>;

interface RegisterFormProps {
  initialData: null;
}

export const RegisterForm: React.FC<RegisterFormProps> = ({ initialData }) => {
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();
  const { toast } = useToast();

  const form = useForm<RegisterFormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: initialData || {},
  });

  const onSubmit = (data: RegisterFormValues) => {
    try {
      setLoading(true);
      myApi.post("/auth/register", data);
      toast({
        title: "Register Success",
        description: "Please login to continue",
      });
      navigate("/auth/login");
    } catch (error) {
      toast({ title: "Register Failed", description: `${error}` });
    } finally {
      setLoading(false);
    }
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
                    disabled={loading}
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
                    disabled={loading}
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
                      disabled={loading}
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
              name="confirmPassword"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-white">Confirm Password</FormLabel>
                  <FormControl>
                    <Input
                      disabled={loading}
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
            disabled={loading}
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
