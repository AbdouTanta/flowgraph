import { Route, Switch } from "wouter";
import Graph from "./components/Graph";
import Home from "./routes/flows";
import { SidebarInset, SidebarProvider } from "./components/ui/sidebar";
import { SiteHeader } from "./components/site-header";
import { AppSidebar } from "./components/app-sidebar";

function App() {
  return (
    <SidebarProvider>
      <AppSidebar variant="inset" />
      <SidebarInset>
        <SiteHeader />
        <Switch>
          <Route path="/" component={Home} />
          <Route path="/graph" component={Graph} />
        </Switch>
      </SidebarInset>
    </SidebarProvider>
  );
}

export default App;
