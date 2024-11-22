"use client";
import React, { useState, useEffect } from "react";
import dynamic from "next/dynamic";
import toast from "react-hot-toast";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
  DropdownMenuShortcut,
} from "@/components/ui/dropdown-menu";
import { Check, ChevronsUpDown } from "lucide-react";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { ChevronDownIcon } from "@heroicons/react/24/outline";

const Editor = dynamic(() => import("@monaco-editor/react"), {
  ssr: false,
  loading: () => (
    <div className="text-white flex items-center justify-center h-[80vh]">
      Loading Editor...
    </div>
  ),
});
const frameworks = [
  {
    value: "next.js",
    label: "Next.js",
  },
  {
    value: "sveltekit",
    label: "SvelteKit",
  },
  {
    value: "nuxt.js",
    label: "Nuxt.js",
  },
  {
    value: "remix",
    label: "Remix",
  },
  {
    value: "astro",
    label: "Astro",
  },
  {
    value: "next.js1",
    label: "Next.js1",
  },
  {
    value: "sveltekit1",
    label: "SvelteKit1",
  },
  {
    value: "nuxt.js1",
    label: "Nuxt.js1",
  },
  {
    value: "remix1",
    label: "Remix1",
  },
  {
    value: "astro1",
    label: "Astro1",
  },
];
const frameworks1 = [
  {
    value: "next.js",
    label: "Next.js",
  },
  {
    value: "sveltekit",
    label: "SvelteKit",
  },
  {
    value: "nuxt.js",
    label: "Nuxt.js",
  },
  {
    value: "remix",
    label: "Remix",
  },
  {
    value: "astro",
    label: "Astro",
  },
  {
    value: "next.js1",
    label: "Next.js1",
  },
  {
    value: "sveltekit1",
    label: "SvelteKit1",
  },
  {
    value: "nuxt.js1",
    label: "Nuxt.js1",
  },
  {
    value: "remix1",
    label: "Remix1",
  },
  {
    value: "astro1",
    label: "Astro1",
  },
];

const Editorsender = () => {
  const [domLoaded, setDomLoaded] = useState(false);
  const [loading, setLoading] = useState(true);
  const [dropdownSelection, setDropdownSelection] = useState("");
  const [requestJSON, setRequestJSON] = useState("");
  const [typeName, setTypeName] = useState("");
  const [popoverOpen1, setPopoverOpen1] = useState(false);
  const [popoverOpen2, setPopoverOpen2] = useState(false);
  const [value1, setValue1] = useState("");
  const [value2, setValue2] = useState("");

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

  const handleSend = async () => {
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
        value1: value1,
        value2: value2,
      };

      console.log("Send Data:", data);

      const response = await fetch(``, {
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
      console.log("Successful", result);
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
            <DropdownMenu className="">
              <DropdownMenuTrigger className="cursor-pointer bg-[#1e1e1e] h-[45px] text-white px-8 py-2 w-48  hover:bg-[#3e3e3e] flex items-center justify-between gap-2 border border-transparent">
                {dropdownSelection || "Select Type"}{" "}
                <ChevronDownIcon className="h-4 w-4" />
              </DropdownMenuTrigger>
              <DropdownMenuContent className="bg-[#1e1e1e] text-white rounded shadow-lg border border-gray-700 w-48 z-50 font-semibold">
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-2 cursor-pointer"
                  onClick={() => {
                    setDropdownSelection("GET");
                    setRequestJSON("");
                  }}
                >
                  GET
                  <DropdownMenuShortcut>⇧⌘G</DropdownMenuShortcut>
                </DropdownMenuItem>
                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-2 cursor-pointer"
                  onClick={() => setDropdownSelection("CREATE")}
                >
                  CREATE
                  <DropdownMenuShortcut>⇧⌘G</DropdownMenuShortcut>
                </DropdownMenuItem>

                <DropdownMenuItem
                  className="hover:bg-gray-700 rounded px-2 py-3 cursor-pointer"
                  onClick={() => {
                    setDropdownSelection("DELETE");
                    setRequestJSON("");
                  }}
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
      <div className="flex pb-8 gap-[30px] px-8">
        <div className="flex items-center justify-center z-30">
          <Popover open={popoverOpen1} onOpenChange={setPopoverOpen1}>
            <PopoverTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={popoverOpen1}
                className="w-[200px] justify-between border-[#6b6b6b]"
              >
                {value1
                  ? frameworks1.find((framework) => framework.value === value1)
                      ?.label
                  : "Select type..."}
                <ChevronsUpDown className="opacity-50" />
              </Button>
            </PopoverTrigger>
            <PopoverContent
              className="w-[200px] p-0 bg-[#1e1e1e] border-[#6b6b6b] z-[9999]"
              side="bottom"
              align="start"
            >
              <Command>
                <CommandInput
                  placeholder="Search type"
                  className="h-9"
                />
                <CommandList>
                  <CommandEmpty>No type found.</CommandEmpty>
                  <CommandGroup>
                    {frameworks1.map((framework) => (
                      <CommandItem
                        key={framework.value}
                        value={framework.value}
                        onSelect={(currentValue) => {
                          setValue1(
                            currentValue === value1 ? "" : currentValue
                          );
                          setPopoverOpen1(false);
                        }}
                      >
                        {framework.label}
                        <Check
                          className={cn(
                            "ml-auto",
                            value1 === framework.value
                              ? "opacity-100"
                              : "opacity-0"
                          )}
                        />
                      </CommandItem>
                    ))}
                  </CommandGroup>
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </div>

        <div className="flex items-center justify-center z-30">
          <Popover open={popoverOpen2} onOpenChange={setPopoverOpen2}>
            <PopoverTrigger asChild>
              <Button
                variant="outline"
                role="combobox"
                aria-expanded={popoverOpen2}
                className="w-[200px] justify-between border-[#6b6b6b]"
              >
                {value2
                  ? frameworks.find((framework) => framework.value === value2)
                      ?.label
                  : "Select type..."}
                <ChevronsUpDown className="opacity-50" />
              </Button>
            </PopoverTrigger>
            <PopoverContent
              className="w-[200px] p-0 bg-[#1e1e1e] border-[#6b6b6b] z-[9999]"
              side="bottom"
              align="start"
            >
              <Command>
                <CommandInput
                  placeholder="Search framework..."
                  className="h-9"
                />
                <CommandList>
                  <CommandEmpty>No type found.</CommandEmpty>
                  <CommandGroup>
                    {frameworks.map((framework) => (
                      <CommandItem
                        key={framework.value}
                        value={framework.value}
                        onSelect={(currentValue) => {
                          setValue2(
                            currentValue === value2 ? "" : currentValue
                          );
                          setPopoverOpen2(false);
                        }}
                      >
                        {framework.label}
                        <Check
                          className={cn(
                            "ml-auto",
                            value2 === framework.value
                              ? "opacity-100"
                              : "opacity-0"
                          )}
                        />
                      </CommandItem>
                    ))}
                  </CommandGroup>
                </CommandList>
              </Command>
            </PopoverContent>
          </Popover>
        </div>
      </div>

      <div className="px-8 bg-[#1e1e1e] border-[#6b6b6b] ">
        {domLoaded && (
          <div
            style={{ height: "50vh" }}
            className="border-[#6b6b6b] border-[1px]"
          >
            <Editor
              height="45vh"
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

export default Editorsender;
