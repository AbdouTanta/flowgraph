import { useLocation } from "wouter";
import { useAuthStore } from "./auth-store";
import { getDecodedJWTFromCookie } from "./jwt";
import { useEffect } from "react";
import { SLUGS } from "./route-slugs";

function AuthProvider({ children }: { children: React.ReactNode }) {
  const [location, navigate] = useLocation();
  const setUser = useAuthStore((state) => state.setUser);

  // TODO - Make a real auth loader that blocks rendering until confirming auth state
  useEffect(() => {
    const isAuthRoute = location === SLUGS.LOGIN || location === SLUGS.REGISTER;

    // Check for authenticated user in cookies
    const decodedToken = getDecodedJWTFromCookie("token");
    if (decodedToken) {
      setUser({
        id: decodedToken.id,
        username: decodedToken.username,
        email: decodedToken.email,
      });
      return;
    }

    // No token found
    if (!isAuthRoute) {
      navigate(SLUGS.LOGIN);
    }
  }, []);

  return children;
}

export { AuthProvider };
