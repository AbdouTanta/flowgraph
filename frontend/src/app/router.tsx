import { Route, Switch } from "wouter";
import Flows from "./routes/flows/flows-list";
import Flow from "./routes/flows/loaded-flow";
import Index from "./routes";
import NewFlow from "./routes/flows/new-flow";
import { LoginPage } from "./routes/auth/login";
import { RegisterPage } from "./routes/auth/register";
import { AppLayout } from "./layout";

// TODO - Handle layout more elegantly
export function AppRouter() {
  return (
    <Switch>
      <Route path="/login" component={LoginPage} />
      <Route path="/register" component={RegisterPage} />
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
