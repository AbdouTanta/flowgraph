import { useQuery } from "@tanstack/react-query";
import type { Flow } from "../types/flows";
import { apiClient } from "@/lib/api-client";

const useFlow = (id?: string) => {
  return useQuery<{ flow: Flow } | null>({
    queryKey: ["flow", id],
    queryFn: async () => {
      if (!id) return null; // If no id, return null
      const response = await apiClient.get(`/api/flows/${id}`);
      if (!response.ok) {
        throw new Error("Failed to fetch flow");
      }
      return response.json();
    },
    enabled: !!id, // Only run this query if id is provided
  });
};

export { useFlow };
