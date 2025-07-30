import { create } from "zustand";

type State = {
  selectedNodeId: string | null;
  selectedFlowId: string | null;
  updateSelectedNodeId: (id: string) => void;
  updateSelectedFlowId: (id: string) => void;
  resetSelectedNodeId: () => void;
  resetSelectedFlowId: () => void;
};

const useCanvas = create<State>((set) => ({
  selectedNodeId: null,
  selectedFlowId: null,
  updateSelectedNodeId: (id: string) => set({ selectedNodeId: id }),
  updateSelectedFlowId: (id: string) => set({ selectedFlowId: id }),
  resetSelectedNodeId: () => set({ selectedNodeId: null }),
  resetSelectedFlowId: () => set({ selectedFlowId: null }),
}));

export default useCanvas;
