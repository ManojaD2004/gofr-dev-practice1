"use client";
import React, { useState } from "react";
import dynamic from "next/dynamic";

const Editor = dynamic(() => import("@monaco-editor/react"), { ssr: false });

const Datasender = () => {
  const [requestJSON, setRequestJSON] = useState(`{
        "name": "John Doe",
        "requestType": "GET",
        "endpoint": "/api/example",
        "headers": {
          "Authorization": "Bearer token"
        }
      }`);
  const [savedContent, setSavedContent] = useState("");

  const handleSave = () => {
    setSavedContent(requestJSON); 
    console.log("Saved Content:", requestJSON);
  };

  const handleEditorChange = (value) => {
    setRequestJSON(value); 
  };

  return (
    <div className="">
      <header className="h-[55px] flex items-center justify-end px-8">
        <button
          className="bg-[#0654a4] px-[20px] font-semibold text-sm py-[10px] text-white rounded"
          onClick={handleSave} 
        >
          Save
        </button>
      </header>
      <Editor
        height="100vh"
        value={requestJSON}
        language="json"
        theme="vs-dark"
        onChange={handleEditorChange} 
      />
    </div>
  );
};

export default Datasender;
