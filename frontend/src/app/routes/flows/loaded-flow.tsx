import ErrorPage from "@/components/error/error-page";
import { useFlow } from "@/features/flows/api/get-flow-by-id";
import Canvas from "@/features/flows/components/canvas";

export default function Flow({ id }: { id?: string }) {
  const loadedFlow = useFlow(id);

  if (loadedFlow.isLoading) {
    return <div>Loading...</div>;
  }

  if (loadedFlow.isError) {
    const message = `Error loading flow: ${loadedFlow.error.message}`;
    return <ErrorPage message={message} />;
  }

  if (!loadedFlow.data) {
    return <div>No flow found.</div>;
  }

  // If the flow is loaded, pass it to the Canvas component
  return <Canvas loadedFlow={loadedFlow.data.flow} />;
}
