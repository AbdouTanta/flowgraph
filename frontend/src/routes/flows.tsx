import type { Flow } from "@/types/flows";
import { useQuery } from "@tanstack/react-query";
import { Link } from "wouter";

// Renders list of flows
export default function Flows() {
  const flows = useQuery<{ flows: Flow[] }>({
    queryKey: ["flows"],
    queryFn: async () => {
      const response = await fetch("/api/flows");
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.json();
    },
  });

  if (flows.isLoading) {
    return <div>Loading...</div>;
  }

  if (flows.isError) {
    return <div>Error loading flows: {flows.error.message}</div>;
  }

  if (!flows.data || flows.data.flows.length === 0) {
    return <div>No flows available.</div>;
  }

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">Your flows</h1>
      <ul className="space-y-2">
        {flows.data.flows.map((flow) => (
          <Link
            href={`/flows/${flow.id}`}
            className="flex flex-row gap-2 cursor-pointer"
            key={flow.id}
          >
            <li className="p-4 border rounded-md hover:bg-gray-100 w-full">
              <h2 className="text-lg font-semibold">{flow.name}</h2>
              <p>{flow.description || "No description"}</p>
            </li>
          </Link>
        ))}
      </ul>
    </div>
  );
}
