import type { Edge, Node } from "@xyflow/react";

export type Flow = {
  id: string;
  name: string;
  description: string;
  nodes: Node[];
  edges: Edge[];
};

export type ICreateFlow = Omit<Flow, "id">;
