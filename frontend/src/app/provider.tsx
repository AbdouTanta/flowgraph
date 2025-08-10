import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { queryConfig } from "@/lib/react-query";

export function AppProvider({ children }: { children: React.ReactNode }) {
  const queryClient = new QueryClient({ defaultOptions: queryConfig });

  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  );
}
