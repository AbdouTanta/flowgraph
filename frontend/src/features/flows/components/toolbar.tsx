import { Button } from "@/components/ui/button";

export default function Toolbar({ onAddNode }: { onAddNode?: () => void }) {
  return (
    <>
      <div className="absolute top-4 right-4 flex gap-2">
        {/* Add Node button */}
        <Button className="z-10 cursor-pointer" onClick={onAddNode}>
          Add Node
        </Button>
      </div>
    </>
  );
}
