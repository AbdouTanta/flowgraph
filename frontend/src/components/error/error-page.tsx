import { IconFaceIdError } from "@tabler/icons-react";

export default function ErrorPage({ message }: { message: string }) {
  return (
    <div className="h-full w-full flex flex-col items-center justify-center">
      <IconFaceIdError className="size-24" />
      <p>{message}</p>
    </div>
  );
}
