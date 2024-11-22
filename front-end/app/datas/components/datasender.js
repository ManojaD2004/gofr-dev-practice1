"use client";
import React, { useState, useEffect } from "react";
import dynamic from "next/dynamic";
import toast, { Toaster } from "react-hot-toast";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
  DropdownMenuShortcut,
} from "@/components/ui/dropdown-menu";
import { ChevronDownIcon } from "@heroicons/react/24/outline";

const Editor = dynamic(() => import("@monaco-editor/react"), {
  ssr: false,
  loading: () => (
    <div className="text-white flex items-center justify-center h-[80vh]">
      Loading Editor...
    </div>
  ),
});

const Datasender = () => {
  const [domLoaded, setDomLoaded] = useState(false);
  const [loading, setLoading] = useState(true);
  const [dropdownSelection, setDropdownSelection] = useState("");
  const [requestJSON, setRequestJSON] = useState(`{
        "name": "John Doe",
        "requestType": "GET",
        "endpoint": "/api/example",
        "headers": {
          "Authorization": "Bearer token"
        }
      }`);
  const [typeName, setTypeName] = useState("");

  useEffect(() => {
    setDomLoaded(true);
  }, []);
  function validateDataType(jsonData) {
    const allowedTypes = ["string", "int", "float32", "float64", "bool"];

    const validateValue = (value, key) => {
      if (typeof value === "string") {
        const seperateWords = value.split(" ");
        const firstWord = seperateWords[0];
        if (!allowedTypes.includes(firstWord)) {
          return {
            isValid: false,
            message: `Invalid data type for key: \`${key}\``,
          };
        }
      } else if (Array.isArray(value)) {
        for (const item of value) {
          if (typeof item === "object" && item !== null) {
            const nestedResult = validateDataType(item);
            if (!nestedResult.isValid) {
              return nestedResult;
            }
          } else {
            return {
              isValid: false,
              message: `Invalid value in array for key: \`${key}\``,
            };
          }
        }
      } else if (typeof value === "object" && value !== null) {
        const nestedResult = validateDataType(value);
        if (!nestedResult.isValid) {
          return nestedResult;
        }
      }
      return { isValid: true };
    };

    for (const [key, value] of Object.entries(jsonData)) {
      const result = validateValue(value, key);
      if (!result.isValid) {
        return result;
      }
    }

    return { isValid: true, message: "All data types are valid." };
  }

  const handleSend = () => {
    try {
      if (!typeName) {
        toast.error("Please enter a type name");
        return;
      }
      if (!dropdownSelection) {
        toast.error("Please select a dropdown option");
        return;
      }
     
      const parsedJSON = JSON.parse(requestJSON);
      const validationResult = validateDataType(parsedJSON);

      if (!validationResult.isValid) {
        toast.error(validationResult.message);
        return;
      }

      const data = {
        input: typeName,
        dropdown: dropdownSelection,
        editorContent: parsedJSON,
      };

      console.log("Send Data:", data);
      toast.success("Data is valid and sent!");
    } catch (error) {
      toast.error("Invalid JSON format.");
      console.error("JSON Parsing Error:", error);
    }
  };

  const handleEditorChange = (value) => {
    setRequestJSON(value);
  };

  return (
    <div className="bg-black min-h-screen ">
      <div className="h-[150px] flex items-center  px-8 text-white justify-between">
        <div className="flex items-center">
          <div className="border-r border-[#6b6b6b]">
            <input
              type="text"
              placeholder="Enter URL or paste text"
              value={typeName}
              onChange={(e) => setTypeName(e.target.value)}
              className="h-[45px] w-[100vh] border-none bg-[#1e1e1e] pl-[20px] font-sans text-sm border-[#6b6b6b] "
            />
          </div>

          <div className="text-white z-30 bg-black font-semibold">
            <DropdownMenu className="">
              <DropdownMenuTrigger className="cursor-pointer bg-[#1e1e1e] h-[45px] text-white px-8 py-2 w-48  hover:bg-[#3e3e3e] flex items-center justify-between gap-2 border border-transparent">
                {dropdownSelection || "Select Type"}{" "}
                <ChevronDownIcon className="h-4 w-4" />
              </DropdownMenuTrigger>
              <DropdownMenuContent className="bg-[#1e1e1e] text-white rounded shadow-lg border border-gray-700 w-48 z-50 font-semibold">
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-2 cursor-pointer"
                  onClick={() => setDropdownSelection("GET")}
                >
                  GET
                  <DropdownMenuShortcut>⇧⌘G</DropdownMenuShortcut>
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-2 cursor-pointer"
                  onClick={() => setDropdownSelection("POST")}
                >
                  POST
                  <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-3 cursor-pointer"
                  onClick={() => setDropdownSelection("DELETE")}
                >
                  DELETE
                  <DropdownMenuShortcut>⇧⌘D</DropdownMenuShortcut>
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-3 cursor-pointer"
                  onClick={() => setDropdownSelection("UPDATE")}
                >
                  UPDATE
                  <DropdownMenuShortcut>⇧⌘U</DropdownMenuShortcut>
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>

        <button
          className="bg-white px-[40px] font-semibold text-sm py-[10px] text-black rounded ml-[40px]"
          onClick={handleSend}
        >
          Send
        </button>
      </div>

      <div className="px-8 bg-black">
        {domLoaded && (
          <div style={{ height: "80vh" }}>
            <Editor
              height="80vh"
              value={requestJSON}
              language="json"
              theme="vs-dark"
              onChange={handleEditorChange}
              onMount={() => setLoading(false)}
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default Datasender;
