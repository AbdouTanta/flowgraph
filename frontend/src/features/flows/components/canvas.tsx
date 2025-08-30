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
import Toolbar from "./toolbar";
import NodeMenu from "./node-menu";
import useCanvas from "@/hooks/use-canvas";
import type { Flow } from "../types/flows";
import { useSidebar } from "@/components/ui/sidebar";

// If flow is provided, it will render it
// If not, the user can create a new flow
export default function Canvas({ loadedFlow }: { loadedFlow: Flow }) {
  const { setOpen } = useSidebar();
  const { selectedNodeId, updateSelectedNodeId, resetSelectedNodeId } =
    useCanvas();
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>([]);

  // If flow data is loaded, set the nodes and edges
  // This effect runs when the flow data is fetched successfully
  useEffect(() => {
    if (loadedFlow) {
      const { nodes, edges } = loadedFlow;
      if (nodes && edges) {
        setNodes(nodes);
        setEdges(edges);
      }
    }
  }, [loadedFlow, setNodes, setEdges]);

  // Close sidebar when flow editor is opened
  useEffect(() => {
    setOpen(false);
    // Leaving this empty dependency array to run only once on mount
  }, []);

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds) => addEdge(params, eds)),
    [setEdges],
  );

  const onNodeClick = useCallback(
    (event: React.MouseEvent, node: Node) => {
      console.log("Node clicked:", event, node);
      // Update the selected node ID in the canvas store
      updateSelectedNodeId(node.id);
    },
    [updateSelectedNodeId],
  );

  const onPaneClick = useCallback(() => {
    // Reset the selected node ID when clicking on the pane
    resetSelectedNodeId();
  }, [resetSelectedNodeId]);

  return (
    <div style={{ width: "100%", height: "100%" }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onNodeClick={onNodeClick}
        onConnect={onConnect}
        onPaneClick={onPaneClick}
        fitView={true}
      >
        <Toolbar
          onAddNode={() => {
            const newNode: Node = {
              id: `${nodes.length + 1}`,
              position: { x: Math.random() * 400, y: Math.random() * 400 },
              data: { label: `Node ${nodes.length + 1}` },
            };
            setNodes((nds) => nds.concat(newNode));
          }}
        />
        <NodeMenu id={selectedNodeId} />
        <Controls />
        <MiniMap />
        <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
      </ReactFlow>
    </div>
  );
}
