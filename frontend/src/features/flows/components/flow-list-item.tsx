import { Link } from "wouter";
import type { Flow } from "../types/flows";

export default function FlowListItem({ flow }: { flow: Flow }) {
  return (
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
  );
}
