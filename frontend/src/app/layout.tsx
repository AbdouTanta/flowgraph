import { SidebarInset, SidebarProvider } from "../components/ui/sidebar";
import { SiteHeader } from "@/components/nav/site-header";
import { AppSidebar } from "@/components/nav/app-sidebar";

export function AppLayout({ children }: { children: React.ReactNode }) {
  return (
    <SidebarProvider>
      <AppSidebar variant="floating" />
      <SidebarInset>
        <SiteHeader />
        {children}
      </SidebarInset>
    </SidebarProvider>
  );
}
