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
import { useRegister } from "../api/register";
import { toast } from "sonner";
import { useLocation } from "wouter";
import { useAuthStore } from "@/lib/auth-store";
import type { User } from "../types/user";
import { SLUGS } from "@/lib/route-slugs";

interface Inputs {
  email: string;
  username: string;
  password: string;
}

export function RegisterForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const methods = useForm<Inputs>({});
  const { handleSubmit } = methods;
  const [, navigate] = useLocation();

  const setUser = useAuthStore((state) => state.setUser);
  const register = useRegister({
    mutationConfig: {
      onSuccess: (res: { token: string; user: User }) => {
        setUser(res.user);
        cookieStore.set("token", res.token);
        navigate("/flows");
      },
      onError: (error: any) => {
        console.error("Register error:", error);
        toast.error("Register failed. Please check your input.");
      },
    },
  });

  const onSubmit = (data: any) => {
    register.mutate({
      email: data.email,
      username: data.username,
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
                <Label htmlFor="username">Username</Label>
                <Input
                  id="username"
                  type="text"
                  required
                  {...methods.register("username")}
                />
              </div>
              <div className="grid gap-3">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
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
                  Register
                </Button>
              </div>
            </div>
            <div className="mt-4 text-center text-sm">
              Already have an account? {""}
              <a href={SLUGS.LOGIN} className="underline underline-offset-4">
                Login
              </a>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
