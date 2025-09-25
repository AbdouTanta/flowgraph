import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useForm } from "react-hook-form";
import { useLogin } from "../api/login";
import { toast } from "sonner";
import { useLocation } from "wouter";
import { useAuthStore } from "@/lib/auth-store";
import type { User } from "../types/user";
import { SLUGS } from "@/lib/route-slugs";

interface Inputs {
  email: string;
  password: string;
}

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const methods = useForm<Inputs>({
    defaultValues: {
      email: "test@test.com",
      password: "test",
    },
  });
  const { handleSubmit } = methods;
  const [, navigate] = useLocation();

  const setUser = useAuthStore((state) => state.setUser);
  const login = useLogin({
    mutationConfig: {
      onSuccess: async (res: { data: { token: string; user: User } }) => {
        setUser(res.data.user);
        await cookieStore.set("token", res.data.token);
        navigate("/flows");
      },
      onError: (error: any) => {
        console.error("Login error:", error);
        toast.error("Login failed. Please check your credentials.");
      },
    },
  });

  const onSubmit = (data: any) => {
    login.mutate({
      email: data.email,
      password: data.password,
    });
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>Login to your account</CardTitle>
          <CardDescription>
            Enter your email below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-3">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  required
                  {...methods.register("email")}
                />
              </div>
              <div className="grid gap-3">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  <a
                    href="#"
                    className="ml-auto inline-block text-sm underline-offset-4 hover:underline"
                  >
                    Forgot your password?
                  </a>
                </div>
                <Input
                  id="password"
                  type="password"
                  required
                  {...methods.register("password")}
                />
              </div>
              <div className="flex flex-row gap-3">
                <Button type="submit" className="w-full">
                  Login
                </Button>
                {/* <Button variant="outline" className="w-full">
                  Login with Google
                </Button> */}
              </div>
            </div>
            <div className="mt-4 text-center text-sm">
              Don&apos;t have an account?{" "}
              <a href={SLUGS.REGISTER} className="underline underline-offset-4">
                Sign up
              </a>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
