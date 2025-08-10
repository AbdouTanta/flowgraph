import { Button } from "@/components/ui/button";

export default function Toolbar({
  onSaveGraph,
  onAddNode,
}: {
  onSaveGraph?: () => void;
  onAddNode?: () => void;
}) {
  return (
    <div className="absolute top-4 right-4 flex gap-2">
      <Button className="z-10 cursor-pointer" onClick={onSaveGraph}>
        Save Graph
      </Button>
      {/* Add Node button */}
      <Button className="z-10 cursor-pointer" onClick={onAddNode}>
        Add Node
      </Button>
    </div>
  );
}
