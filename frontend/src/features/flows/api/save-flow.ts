import { useMutation } from "@tanstack/react-query";
import type { ICreateFlow } from "../types/flows";

const useSaveFlow = () => {
  return useMutation({
    mutationFn: async (graphData: ICreateFlow) => {
      const response = await fetch("/api/flows", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(graphData),
      });
      if (!response.ok) {
        throw new Error("Failed to save graph");
      }
      return response.json();
    },
  });
};

export { useSaveFlow };
