import { Route, Switch } from "wouter";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import Flows from "./routes/flows";
import { SidebarInset, SidebarProvider } from "./components/ui/sidebar";
import { SiteHeader } from "./components/nav/site-header";
import { AppSidebar } from "./components/nav/app-sidebar";
import Flow from "./routes/flow";
import Index from "./routes";
import NewFlow from "./routes/new-flow";
import { Toaster } from "./components/ui/sonner";
import { queryConfig } from "./lib/react-query";

function Wrapper({ children }: { children: React.ReactNode }) {
  // To be used for layout, hooks etc.
  return children;
}

function App() {
  const queryClient = new QueryClient({ defaultOptions: queryConfig });

  return (
    <QueryClientProvider client={queryClient}>
      <SidebarProvider>
        <Wrapper>
          <AppSidebar variant="floating" />
          <SidebarInset>
            <SiteHeader />
            <Switch>
              <Route path="/" component={Index} />
              <Route path="/flows" component={Flows} />
              <Route path="/flows/new" component={NewFlow} />
              <Route path="/flows/:id">
                {({ id }) => <Flow id={id as string} />}
              </Route>
            </Switch>
          </SidebarInset>
          <Toaster />
        </Wrapper>
      </SidebarProvider>
    </QueryClientProvider>
  );
}

export default App;
