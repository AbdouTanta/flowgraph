import ErrorPage from "@/components/error/error-page";
import { useFlows } from "@/features/flows/api/get-flows";
import FlowListItem from "@/features/flows/components/flow-list-item";

export default function Flows() {
  const flowsQuery = useFlows();

  if (flowsQuery.isLoading) {
    return <div>Loading...</div>;
  }

  if (flowsQuery.isError) {
    const message = `Error loading flows: ${flowsQuery.error.message}`;
    return <ErrorPage message={message} />;
  }

  const flowsList = flowsQuery.data?.data?.flows;

  if (!flowsList || flowsList?.length === 0) {
    return <div className="p-4">No flows available.</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">Your flows</h1>
      <ul className="space-y-2">
        {flowsList.map((flow) => (
          <FlowListItem key={flow.id} flow={flow} />
        ))}
      </ul>
    </div>
  );
}
