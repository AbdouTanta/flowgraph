import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { queryConfig } from "@/config/react-query";
import { AuthProvider } from "@/lib/auth-provider";
import { Toaster } from "sonner";

export function AppProvider({ children }: { children: React.ReactNode }) {
  const queryClient = new QueryClient({ defaultOptions: queryConfig });

  return (
    <QueryClientProvider client={queryClient}>
      <Toaster />
      <AuthProvider>{children}</AuthProvider>
    </QueryClientProvider>
  );
}
