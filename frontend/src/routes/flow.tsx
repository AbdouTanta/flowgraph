import FlowEditor from "@/components/editor/canvas";

export default function Flow({ id }: { id?: string }) {
  return <FlowEditor id={id} />;
}
