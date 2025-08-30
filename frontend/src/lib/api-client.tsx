import { getCookie } from "./jwt";
import { SLUGS } from "./route-slugs";

// Helpers for making API requests with auth token handling
export const apiClient = {
  async request(url: string, options: RequestInit = {}) {
    const token = getCookie("token");

    const config: RequestInit = {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...(token && { Authorization: `Bearer ${token}` }),
        ...options.headers,
      },
    };

    const response = await fetch(url, config);

    // Handle auth errors globally
    if (response.status === 401) {
      // Clear auth state and redirect to login
      cookieStore.delete("token");
      // You might want to clear your auth store here too
      window.location.href = SLUGS.LOGIN;
    }

    return response;
  },

  get: (url: string) => apiClient.request(url),
  post: (url: string, data: any) =>
    apiClient.request(url, {
      method: "POST",
      body: JSON.stringify(data),
    }),
  put: (url: string, data: any) =>
    apiClient.request(url, {
      method: "PUT",
      body: JSON.stringify(data),
    }),
  delete: (url: string) => apiClient.request(url, { method: "DELETE" }),
};
