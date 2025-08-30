import { IconCirclePlusFilled, type Icon } from "@tabler/icons-react";

import {
  SidebarGroup,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import { Link } from "wouter";
import { CreateFlowModal } from "@/features/flows/components/create-flow-modal";
import { useState } from "react";
import { useCreateFlow } from "@/features/flows/api/create-flow";
import { toast } from "sonner";

export function NavMain({
  items,
}: {
  items: {
    title: string;
    url: string;
    icon?: Icon;
  }[];
}) {
  const [open, setOpen] = useState(false);
  const createFlow = useCreateFlow();

  return (
    <SidebarGroup>
      <SidebarGroupContent className="flex flex-col gap-2">
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              tooltip="Create new flow"
              onClick={() => setOpen(true)}
              className="bg-primary text-primary-foreground hover:bg-primary/90 hover:text-primary-foreground active:bg-primary/90 active:text-primary-foreground min-w-8 duration-200 ease-linear"
            >
              <IconCirclePlusFilled />
              <span>Create new flow</span>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
        <SidebarMenu>
          {items.map((item) => (
            <SidebarMenuItem key={item.title}>
              <Link href={item.url}>
                <SidebarMenuButton tooltip={item.title}>
                  {item.icon && <item.icon />}
                  <span>{item.title}</span>
                </SidebarMenuButton>
              </Link>
            </SidebarMenuItem>
          ))}
        </SidebarMenu>
      </SidebarGroupContent>

      {/* Mount modals */}
      <CreateFlowModal
        open={open}
        onOpenChange={setOpen}
        onSubmit={(name: string, description: string) => {
          createFlow.mutate(
            { name, description },
            {
              onSuccess: () => {
                toast.success("Flow created");
              },
              onError: () => {
                toast.error("Failed to create flow");
              },
            },
          );
        }}
      />
    </SidebarGroup>
  );
}
