import React, { useState, useEffect } from "react";
import { Handle, Position, useReactFlow } from "reactflow";

export default function Databaseoptions({
  data: { name, code, inputValue = "" }, // Default to empty string
  id,
  onInputChange,
}) {
  const { setNodes } = useReactFlow();
  const [localInputValue, setLocalInputValue] = useState(inputValue);

  const handleInputChange = (event) => {
    const value = event.target.value;
    setLocalInputValue(value); // Update local state
    onInputChange(id, value); // Pass input value back to parent
  };

  useEffect(() => {
    // Ensure component stays in sync with parent state
    setLocalInputValue(inputValue || ""); // Ensure it's a non-undefined value
  }, [inputValue]);

  return (
    <div className="flex items-center bg-white p-2 rounded-2xl gap-2 w-[400px]">
      <div>
        <div className="flex items-center">
          <div className="flex flex-grow">
            <p className="text-sm mt-[-2px] text-black ml-[10px]">{name}</p>
          </div>
          <button
            aria-label="Delete Payment Provider"
            className="text-red-500 bg-transparent hover:text-red-700 focus:outline-none"
            onClick={() =>
              setNodes((prevNodes) =>
                prevNodes.filter((node) => node.id !== id)
              )
            }
          >
            ✖
          </button>
        </div>
        <div className="my-[10px]">
          <input
            type="text"
            value={localInputValue || ""} // Ensure it's always a string (even if inputValue is null/undefined)
            onChange={handleInputChange}
            placeholder="Enter value"
            className="border border-gray-300 p-1 rounded-md text-sm w-[360px] text-black ml-[10px]"
          />
        </div>
      </div>
      <Handle
        type="target"
        position={Position.Left}
        className="w-2 h-2 bg-blue-500"
      />
      <Handle
        type="source"
        position={Position.Right}
        className="w-2 h-2 bg-red-500"
      />
    </div>
  );
}
