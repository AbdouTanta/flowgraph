import type { MutationConfig } from "@/config/react-query";
import { useMutation } from "@tanstack/react-query";

interface IRegisterPayload {
  email: string;
  username: string;
  password: string;
}

async function registerUser(payload: IRegisterPayload) {
  return fetch("/api/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(payload),
  }).then((res) => {
    if (!res.ok) {
      throw new Error("Network response was not ok");
    }
    return res.json();
  });
}

type UseRegisterOptions = {
  mutationConfig?: MutationConfig<typeof registerUser>;
};

const useRegister = ({ mutationConfig }: UseRegisterOptions = {}) => {
  const { ...restConfig } = mutationConfig || {};

  return useMutation({
    ...restConfig,
    mutationFn: registerUser,
  });
};

export { useRegister };
