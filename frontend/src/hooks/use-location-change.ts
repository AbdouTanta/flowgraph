// hooks/useSidebarRoute.js
import { useSidebar } from "@/components/ui/sidebar";
import { useEffect } from "react";
import { useLocation } from "wouter";

const useLocationChange = () => {
  const [location] = useLocation();
  const { setOpen } = useSidebar();

  useEffect(() => {
    // If location changes to flows/:id, close the sidebar
    console.log("Location changed to:", location);
    if (location.startsWith("/flows/")) {
      setOpen(false);
    }
  }, [location, setOpen]);
};

export default useLocationChange;
