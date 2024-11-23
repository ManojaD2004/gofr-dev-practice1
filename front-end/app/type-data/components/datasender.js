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
  const [requestJSON, setRequestJSON] = useState("");
  const [typeName, setTypeName] = useState("");
  const [allowedTypes, setAllowedTypes] = useState([
    "string",
    "int",
    "float32",
    "float64",
    "bool",
  ]);
  const [isEditorDisabled, setIsEditorDisabled] = useState(false);

  useEffect(() => {
    setDomLoaded(true);
    typeloading();
  }, []);

  const typeloading = async () => {
    try {
      const response = await fetch(
        `http://localhost:8000/.__gofr__/get-all-types`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();
      console.log("API Response:", result);

      const { data } = result;

      if (!data) {
        throw new Error("Data field not found in the response.");
      }

      const frameworks = Object.keys(data).map((type) => type);

      setAllowedTypes((prevFramework) => Array.from(new Set([...prevFramework, ...frameworks])));

      setLoading(false);
    } catch (error) {
      console.error(error);
    }
  };

  function validateDataType(jsonData) {
    const validateValue = (value, key) => {
      console.log(key, value)
      if (typeof value === "string") {
        const separateWords = value.split(" ");
        const firstWord = separateWords[0];
        // console.log(separateWords);
        if (!allowedTypes.includes(firstWord)) {
          return {
            isValid: false,
            message: `Invalid data type for key: \`${key}\``,
          };
        }
      } else if (Array.isArray(value)) {
        for (const item of value) {
          const nestedResult = validateValue(item, key);
          if (!nestedResult.isValid) {
            return nestedResult;
          }
        } 
      } else if (typeof value === "object" && value !== null) {
        const nestedResult = validateValue(value, key);
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

  const handleSend = async () => {
    try {
      console.log(allowedTypes);
      if (!typeName) {
        toast.error("Please enter a type name");
        return;
      }
      if (!dropdownSelection) {
        toast.error("Please select a dropdown option");
        return;
      }

      let data = {
        typeName: typeName,
      };

      if (dropdownSelection === "CREATE" || dropdownSelection === "UPDATE") {
        if (!requestJSON) {
          toast.error("Editor content cannot be empty for this action.");
          return;
        }

        try {
          const parsedJSON = JSON.parse(requestJSON);
         {/*} const validationResult = validateDataType(parsedJSON);

          if (!validationResult.isValid) {
            toast.error(validationResult.message);
            return;
          } */}

          data.typeBody = parsedJSON;
        } catch (parseError) {
          toast.error("Invalid JSON format in the editor.");
          console.error("JSON Parsing Error:", parseError);
          return;
        }
      }

      console.log("Send Data:", data);

      let apiEndpoint = "";
      if (dropdownSelection === "GET") {
        apiEndpoint = "http://localhost:8000/.__gofr__/get-type";
      } else if (dropdownSelection === "CREATE") {
        apiEndpoint = "http://localhost:8000/.__gofr__/create-type";
      } else if (dropdownSelection === "DELETE") {
        apiEndpoint = "http://localhost:8000/.__gofr__/delete-type";
      } else if (dropdownSelection === "UPDATE") {
        apiEndpoint = "http://localhost:8000/.__gofr__/update-type";
      }

      const response = await fetch(apiEndpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        toast.error("Error creating API");
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const result = await response.json();
      console.log("Successful:", result);
      setRequestJSON(JSON.stringify(result, null, 2));
      typeloading();
      toast.success("Data is valid and sent!");
    } catch (error) {
      toast.error("An error occurred. Please check your input.");
      console.error("Error:", error);
    }
  };

  const handleEditorChange = (value) => {
    setRequestJSON(value);
  };

  const handleDropdownChange = (selection) => {
    setDropdownSelection(selection);

    if (selection === "GET" || selection === "DELETE") {
      setIsEditorDisabled(true);
    } else {
      setIsEditorDisabled(false);
    }

    setRequestJSON("");
  };

  return (
    <div className="bg-[#1e1e1e] min-h-screen ">
      <div className="h-[150px] flex items-center  px-8 text-white justify-between">
        <div className="flex items-center border border-[#6b6b6b] ">
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
            <DropdownMenu>
              <DropdownMenuTrigger className="cursor-pointer bg-[#1e1e1e] h-[45px] text-white px-8 py-2 w-48  hover:bg-[#3e3e3e] flex items-center justify-between gap-2 border border-transparent">
                {dropdownSelection || "Select Type"}{" "}
                <ChevronDownIcon className="h-4 w-4" />
              </DropdownMenuTrigger>
              <DropdownMenuContent className="bg-[#1e1e1e] text-white rounded shadow-lg border border-gray-700 w-48 z-50 font-semibold">
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-2 cursor-pointer"
                  onClick={() => handleDropdownChange("GET")}
                >
                  GET
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-2 cursor-pointer"
                  onClick={() => handleDropdownChange("CREATE")}
                >
                  CREATE
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-3 cursor-pointer"
                  onClick={() => handleDropdownChange("DELETE")}
                >
                  DELETE
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-3 cursor-pointer"
                  onClick={() => handleDropdownChange("UPDATE")}
                >
                  UPDATE
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

      <div className="px-8 bg-[#1e1e1e] border-[#6b6b6b] ">
        {domLoaded && (
          <div
            style={{ height: "52vh" }}
            className="border-[#6b6b6b] border-[1px]"
          >
            <Editor
              height="50vh"
              value={requestJSON}
              language="json"
              theme="vs-dark"
              onChange={handleEditorChange}
              onMount={() => setLoading(false)}
              options={{ readOnly: isEditorDisabled }}
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default Datasender;
