import { Route, Switch } from "wouter";
import Flows from "./routes/flows/flows-list";
import Flow from "./routes/flows/loaded-flow";
import Index from "./routes";
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
        <Switch>
          <Route path="/flows/:id">
            {({ id }) => <Flow id={id as string} />}
          </Route>
          <Route path="/flows" component={Flows} />
        </Switch>
      </AppLayout>
    </Switch>
  );
}
