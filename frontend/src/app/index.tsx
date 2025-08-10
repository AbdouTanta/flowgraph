import { AppLayout } from "./layout";
import { AppProvider } from "./provider";
import { AppRouter } from "./router";

export function App() {
  return (
    <AppProvider>
      <AppLayout>
        <AppRouter />
      </AppLayout>
    </AppProvider>
  );
}
