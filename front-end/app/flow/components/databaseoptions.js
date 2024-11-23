import React from "react";
import { Handle, Position, useReactFlow } from "reactflow";

export default function Databaseoptions({ data: { name, code }, id }) {
  const { setNodes } = useReactFlow();

  return (
    <div className="flex items-center  bg-white p-2 rounded-2xl gap-2 w-36">
      <div className="h-4 w-4">X</div>
      <div className="flex-grow">
        <p className="text-sm mt-[-2px] text-black">{name}</p>
      </div>
      <button
        aria-label="Delete Payment Provider"
        className="text-red-500 bg-transparent hover:text-red-700 focus:outline-none"
        onClick={() =>
          setNodes((prevNodes) => prevNodes.filter((node) => node.id !== id))
        }
      >
        ✖
      </button>
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
