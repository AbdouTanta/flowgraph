// Renders list of flows
export default function Flows() {
  const flows = [
    { id: "1", name: "Flow 1", description: "This is the first flow." },
    { id: "2", name: "Flow 2", description: "This is the second flow." },
    { id: "3", name: "Flow 3", description: "This is the third flow." },
    { id: "4", name: "Flow 4", description: "This is the fourth flow." },
    { id: "5", name: "Flow 5", description: "This is the fifth flow." },
    { id: "6", name: "Flow 6", description: "This is the sixth flow." },
    { id: "7", name: "Flow 7", description: "This is the seventh flow." },
    { id: "8", name: "Flow 8", description: "This is the eighth flow." },
    { id: "9", name: "Flow 9", description: "This is the ninth flow." },
    { id: "10", name: "Flow 10", description: "This is the tenth flow." },
  ];

  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4">Your flows</h1>
      <ul className="space-y-2">
        {flows.map((flow) => (
          <li key={flow.id} className="p-4 border rounded-md hover:bg-gray-100">
            <h2 className="text-lg font-semibold">{flow.name}</h2>
            <p>{flow.description}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}
