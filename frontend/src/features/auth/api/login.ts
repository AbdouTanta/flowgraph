import { useMutation } from "@tanstack/react-query";

interface ILoginPayload {
  email: string;
  password: string;
}

const useLogin = () => {
  return useMutation({
    mutationFn: async (loginPayload: ILoginPayload) => {
      const response = await fetch("/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginPayload),
      });
      if (!response.ok) {
        throw new Error("Failed to login");
      }
      return response.json();
    },
  });
};

export { useLogin };
