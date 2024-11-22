"use client";
import React, { useState,useEffect } from "react";

import toast from "react-hot-toast";
import dynamic from "next/dynamic";
const Editor = dynamic(() => import("@monaco-editor/react"), { ssr: false });

const Page = () => {
    const [domLoaded, setDomLoaded] = useState(false);
    useEffect(() => {
        setDomLoaded(true);
      }, []);
  const [selectedOption, setSelectedOption] = useState("GET");
  const [typeName, setInputValue] = useState("");
  const [requestJSON, setRequestJSON] = useState(`{
    "name": "John Doe",
    "requestType": "GET",
    "endpoint": "/api/example",
    "headers": {
      "Authorization": "Bearer token"
    }
  }`);
  const [requestJSON1, setRequestJSON1] = useState("");

  const [responseJSON, setResponseJSON] = useState("");

  const handleChange = (event) => {
    setSelectedOption(event.target.value);
  };

  const handleChange1 = (event) => {
    setInputValue(event.target.value);
  };

  const handleEditorChange = (value, event) => {
    setRequestJSON(value);
  };
  const handleEditorChange1 = (value, event) => {
    setRequestJSON1(value);
  };

  const handleClick = async () => {
    try {
      const response = await apireq();

      setResponseJSON(JSON.stringify(response, null, 2));
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  const apireq = async () => {
    const reqbody = validateJSON(requestJSON);
    const resbody = validateJSON(requestJSON1);
  
    const data = {
      typeName,
      reqbody,
      resbody,
    };
  
    console.log(data);
  
    try {
      const response = await fetch(
        `https://31f8-115-99-94-143.ngrok-free.app/create-route`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        }
      );
  
      if (!response.ok) {
        toast.error("Error creating API");
        throw new Error(`HTTP error! status: ${response.status}`);
      }
  
      const result = await response.json();
      console.log("Successful", result);
  
      if (result.data === "Hello Tiger!") {
        console.log("Done");
        toast.success("New API created");
      }
  
      setResponseJSON(JSON.stringify(result, null, 2));
    } catch (error) {
      toast.error("Error creating API");
      console.error("Error posting data:", error);
      setResponseJSON(
        "Invalid JSON response or server error. Check console logs."
      );
    }
  };
  

  function handleEditorDidMount(editor, monaco) {
    console.log("Editor Instance:", editor);
    console.log("Monaco Instance:", monaco);
  }

  function handleEditorWillMount(monaco) {
    console.log("Preparing Monaco Editor Instance:", monaco);
  }

  function handleEditorValidation(markers) {
    markers.forEach((marker) =>
      console.log("Validation Issue:", marker.message)
    );
  }

  return (
    <div className="h-full flex flex-col bg-[#1e1e1e] pb-[50px] pt-4">
      <div className="flex ">
        <div className="flex gap-1 border-gray-700 p-[5px] border-[1px] rounded-sm m-[20px] divide-x">
          <div className="border-none ">
            <select
              value={selectedOption}
              onChange={handleChange}
              className="h-[35px] border-none bg-[#1e1e1e] font-bold px-[15px] "
            >
              <option value="get">GET</option>
              <option value="Post">POST</option>
              <option value="Put">PUT</option>
              <option value="Delete">DELETE</option>
              <option value="patch">PATCH</option>
            </select>
          </div>
          <vr className="border-gray-600"></vr>
          <div className="border-none flex-1">
            <input
              type="text"
              placeholder="Enter URL or paste text"
              value={typeName}
              onChange={handleChange1}
              className="h-[35px] w-[160vh] border-none bg-[#1e1e1e] pl-[20px] font-sans text-sm "
            />
          </div>
        </div>
        <div className="m-[20px] ">
          <button
            onClick={handleClick}
            className="bg-blue-500 text-white p-2 rounded h-[48px] font-semibold px-[30px]"
          >
            Send
          </button>
        </div>
      </div>

      <div className="w-[full] mx-[20px] border-[1px] border-gray-700 rounded-sm">
        {domLoaded &&(
        <Editor
          height="38vh"
          value={requestJSON}
          language="json"
          theme="vs-dark"
          onChange={handleEditorChange}
          onMount={handleEditorDidMount}
          beforeMount={handleEditorWillMount}
          onValidate={handleEditorValidation}
        
        />
        )}
      </div>

      <div className="font-semibold py-[10px] text-sm text-white ml-[42px]">
        Response
      </div>
      <div className="w-[full] mx-[20px] border-[1px] border-gray-700 rounded-sm">
        <Editor
          height="38vh"
          value={responseJSON}
          language="json"
          theme="vs-dark"
          onChange={handleEditorChange1}
          onMount={handleEditorDidMount}
          beforeMount={handleEditorWillMount}
          onValidate={handleEditorValidation}
        />
      </div>
    </div>
  );
};

export default Page;
