import ErrorPage from "@/components/error/error-page";
import { useFlows } from "@/features/flows/api/get-flows";
import FlowListItem from "@/features/flows/components/flow-list-item";

export default function Flows() {
  const flows = useFlows();

  if (flows.isLoading) {
    return <div>Loading...</div>;
  }

  if (flows.isError) {
    const message = `Error loading flows: ${flows.error.message}`;
    return <ErrorPage message={message} />;
  }

  if (!flows.data || flows.data.flows.length === 0) {
    return <div>No flows available.</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">Your flows</h1>
      <ul className="space-y-2">
        {flows.data.flows.map((flow) => (
          <FlowListItem key={flow.id} flow={flow} />
        ))}
      </ul>
    </div>
  );
}
