import { Separator } from "@/components/ui/separator";
import { LoginForm } from "./components/login-form";

export default function LoginPage() {
  return (
    <div className="flex h-screen">
      <div className="hidden lg:flex items-center justify-center flex-1 text-black">
        <div className="max-w-md text-center">
          <h1 className="text-3xl font-bold">Login</h1>
          <p className="mt-4 text-lg">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua.
          </p>
        </div>
      </div>
      <Separator orientation="vertical" />
      <div className="w-full lg:w-1/2 flex items-center justify-center">
        <div className="max-w-md w-full p-6">
          <h1 className="text-3xl font-semibold mb-4 text-center uppercase">
            Masuk
          </h1>
          <h3 className="text-center text-sm">Masuk ke akun Anda</h3>
          <LoginForm initialData={null} />
        </div>
      </div>
    </div>
  );
}
