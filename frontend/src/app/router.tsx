import { Route, Switch } from "wouter";
import Flows from "./routes/flows/flows-list";
import Flow from "./routes/flows/loaded-flow";
import Index from "./routes";
import NewFlow from "./routes/flows/new-flow";

export function AppRouter() {
  return (
    <Switch>
      <Route path="/" component={Index} />
      <Route path="/flows" component={Flows} />
      <Route path="/flows/new" component={NewFlow} />
      <Route path="/flows/:id">{({ id }) => <Flow id={id as string} />}</Route>
    </Switch>
  );
}
