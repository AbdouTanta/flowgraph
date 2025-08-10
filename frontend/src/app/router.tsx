import { Route, Switch } from "wouter";
import Flows from "./routes/flows/flows-list";
import Flow from "./routes/flows/loaded-flow";
import Index from "./routes";
import NewFlow from "./routes/flows/new-flow";
import Login from "./routes/auth/login";
import { AppLayout } from "./layout";

// TODO - Handle layout more elegantly
export function AppRouter() {
  return (
    <Switch>
      <Route path="/auth/login" component={Login} />
      <Route path="/" component={Index} />
      <AppLayout>
        <Route path="/flows" component={Flows} />
        <Route path="/flows/new" component={NewFlow} />
        <Route path="/flows/:id">
          {({ id }) => <Flow id={id as string} />}
        </Route>
      </AppLayout>
    </Switch>
  );
}
