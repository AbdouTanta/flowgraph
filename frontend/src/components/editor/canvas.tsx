import { useCallback, useEffect } from "react";
import {
  ReactFlow,
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  addEdge,
  BackgroundVariant,
  type Connection,
  type Node,
  type Edge,
} from "@xyflow/react";

import "@xyflow/react/dist/style.css";
import { useMutation, useQuery } from "@tanstack/react-query";
import type { Flow, ICreateFlow } from "@/types/flows";
import Toolbar from "./toolbar";
import { toast } from "sonner";
import NodeMenu from "./node-menu";
import { useSidebar } from "../ui/sidebar";

const initialNodes: Node[] = [
  {
    id: "1",
    type: "input",
    position: { x: 0, y: 0 },
    data: { label: "Input Node" },
  },
  {
    id: "2",
    position: { x: 100, y: 100 },
    data: { label: "Default Node" },
  },
];
const initialEdges: Edge[] = [
  {
    id: "e1-2",
    source: "1",
    target: "2",
  },
];

// If id is provided, it will load the flow with that id
// If not, it will create a new empty flow
export default function Canvas({ id }: { id?: string }) {
  const { setOpen } = useSidebar();
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>(initialEdges);

  // If id is provided, fetch the flow data from the server
  const flow = useQuery<{ flow: Flow } | null>({
    queryKey: ["flow", id],
    queryFn: async () => {
      if (!id) return null; // If no id, return null
      const response = await fetch(`/api/flows/${id}`);
      if (!response.ok) {
        throw new Error("Failed to fetch flow");
      }
      return response.json();
    },
    enabled: !!id, // Only run this query if id is provided
  });

  // If flow data is loaded, set the nodes and edges
  // This effect runs when the flow data is fetched successfully
  useEffect(() => {
    if (flow.data) {
      const { nodes: flowNodes, edges: flowEdges } = flow.data.flow;
      if (flowNodes && flowEdges) {
        setNodes(flowNodes);
        setEdges(flowEdges);
      }
    }
  }, [flow.data, setNodes, setEdges]);

  // Close sidebar when flow editor is opened
  useEffect(() => {
    setOpen(false);
  }, []);

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds) => addEdge(params, eds)),
    [setEdges],
  );

  const saveFlow = useMutation({
    mutationFn: async (graphData: ICreateFlow) => {
      const response = await fetch("/api/flows", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(graphData),
      });
      if (!response.ok) {
        throw new Error("Failed to save graph");
      }
      return response.json();
    },
  });

  if (flow.isLoading) {
    return <div>Loading flow...</div>;
  }

  if (flow.isError) {
    return <div>Error loading flow: {flow.error.message}</div>;
  }

  return (
    <div style={{ width: "100%", height: "100%" }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
        fitView={true}
      >
        <Toolbar
          onSaveGraph={() => {
            const graphData = { nodes, edges };
            // Here you can implement saving logic, e.g., sending to a server
            saveFlow.mutate(
              {
                // TODO - replace with actual graph data
                name: "lol",
                description: "even more lolz",
                ...graphData,
              },
              {
                onSuccess: () => {
                  toast("Graph saved successfully!");
                },
                onError: (error) => {
                  console.error("Error saving graph:", error);
                  toast("Failed to save graph");
                },
              },
            );
          }}
          onAddNode={() => {
            const newNode: Node = {
              id: `${nodes.length + 1}`,
              position: { x: Math.random() * 400, y: Math.random() * 400 },
              data: { label: `Node ${nodes.length + 1}` },
            };
            setNodes((nds) => nds.concat(newNode));
          }}
        />
        <NodeMenu />
        <Controls />
        <MiniMap />
        <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
      </ReactFlow>
    </div>
  );
}
