import { useMutation } from "@tanstack/react-query";
import type { ICreateFlow } from "../types/flows";
import { apiClient } from "@/lib/api-client";

const useCreateFlow = () => {
  return useMutation({
    mutationFn: async (payload: ICreateFlow) => {
      const response = await apiClient.post("/api/flows", payload);
      if (!response.ok) {
        throw new Error("Failed to create flow");
      }
      return response.json();
    },
  });
};

export { useCreateFlow };
