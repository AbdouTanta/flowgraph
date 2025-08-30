import { useQuery } from "@tanstack/react-query";
import type { Flow } from "../types/flows";
import { apiClient } from "@/lib/api-client";

export function useFlows() {
  return useQuery<{ flows: Flow[] }>({
    queryKey: ["flows"],
    queryFn: async () => {
      const response = await apiClient.get("/api/flows");
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    },
  });
}
