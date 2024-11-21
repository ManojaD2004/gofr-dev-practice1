"use client";
import React, { useState } from "react";
import Editor from "@monaco-editor/react";

const Page = () => {
  const [selectedOption, setSelectedOption] = useState("");
  const [inputValue, setInputValue] = useState("");
  const [requestJSON, setRequestJSON] = useState(`{
    "name": "John Doe",
    "requestType": "GET",
    "endpoint": "/api/example",
    "headers": {
      "Authorization": "Bearer token"
    }
  }`);
  const [requestJSON1,setRequestJSON1]=useState("")

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
  const handleEditorChange1=(value,event)=>{
    setRequestJSON1(value);
  }

  
  const handleClick = async () => {
    console.log("Selected Option:", selectedOption);
    console.log("Input Value:", inputValue);
    console.log("Editor Content:", requestJSON);
    console.log("Editor1 Content:",requestJSON1);

    try {
  
      const response = await apireq();

     
      setResponseJSON(JSON.stringify(response, null, 2)); 
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  const apireq = async () => {
    console.log("hello");
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
    <div className="h-full flex flex-col bg-[#1e1e1e] pb-[50px]">
      <div className="flex ">
        <div className="flex gap-1 border-gray-700 p-[5px] border-[1px] rounded-sm m-[20px] divide-x">
          <div className="border-none ">
            <select
              value={selectedOption}
              onChange={handleChange}
              className="h-[35px] border-none bg-[#1e1e1e] font-bold px-[15px]"
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
              value={inputValue}
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
      </div>

      <div className="font-semibold py-[10px] text-sm text-white ml-[20px]">Response</div>
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
