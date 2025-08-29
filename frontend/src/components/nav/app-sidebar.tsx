import * as React from "react";
import {
  IconFolder,
  IconInnerShadowTop,
  IconSettings,
} from "@tabler/icons-react";

import { NavMain } from "@/components/nav/nav-main";
import { NavBottom } from "@/components/nav/nav-bottom";
import { NavUser } from "@/components/nav/nav-user";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Link, useLocation } from "wouter";
import { useAuthStore } from "@/lib/auth-store";
import { SLUGS } from "@/lib/route-slugs";
import { setCookie } from "@/lib/jwt";

const data = {
  user: {
    name: "abdoutanta",
    email: "at@example.com",
    avatar: "/avatars/shadcn.jpg",
  },
  navMain: [
    {
      title: "Flows",
      url: "/",
      icon: IconFolder,
    },
  ],
  navBottom: [
    {
      title: "Settings",
      url: "/settings",
      icon: IconSettings,
    },
  ],
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const user = useAuthStore((state) => state.user);
  const [, navigate] = useLocation();

  if (!user) return null;

  return (
    <Sidebar collapsible="offcanvas" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              asChild
              className="data-[slot=sidebar-menu-button]:!p-1.5"
            >
              <Link href="/">
                <IconInnerShadowTop className="!size-5" />
                <span className="text-base font-semibold">Flowgraph.</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
        <NavBottom items={data.navBottom} className="mt-auto" />
      </SidebarContent>
      <SidebarFooter>
        <NavUser
          user={user!}
          // TODO - Put elsewhere and manage login and logout properly
          onLogout={() => {
            useAuthStore.getState().setUser(null);
            setCookie("token", "");
            navigate(SLUGS.LOGIN);
          }}
        />
      </SidebarFooter>
    </Sidebar>
  );
}
