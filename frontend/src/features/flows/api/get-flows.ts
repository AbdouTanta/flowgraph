import { useQuery } from "@tanstack/react-query";
import type { Flow } from "../types/flows";
import { apiClient } from "@/lib/api-client";

export function useFlows() {
  return useQuery<{ data: { flows: Flow[] } }>({
    queryKey: ["flows"],
    queryFn: async () => {
      const response = await apiClient.get("/api/flows");
      if (!response.ok) {
        const errorResponse = await response.json();
        if (errorResponse && errorResponse.message) {
          throw new Error(errorResponse.message);
        }
        throw new Error(`Error fetching flows: ${response.statusText}`);
      }
      return response.json();
    },
  });
}
