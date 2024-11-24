"use client";
import React, { useState } from "react";
import toast from "react-hot-toast";

const CsvUploader = () => {
  const [file, setFile] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [tablename, setTypeName] = useState("");

  const handleFileChange = (event) => {
    const uploadedFile = event.target.files[0];
    if (uploadedFile && uploadedFile.type === "text/csv") {
      setFile(uploadedFile);
      setErrorMessage("");
    } else {
      setErrorMessage("Please upload a valid CSV file.");
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (!file) {
      toast.error("Please upload a file.");
      return;
    }

    if (!tablename) {
      toast.error("Please enter a table name.");
      return;
    }

    const formData = new FormData();
    formData.append("tablename", tablename); // Match Go's "tablename"
    formData.append("filename", file.name); // Match Go's "filename"
    formData.append("csvfile", file); // Match Go's "csvfile"

    try {
      setUploading(true);
      setErrorMessage("");

      const response = await fetch(`https://a0e4-119-82-122-154.ngrok-free.app/convertcsv`, {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const result = await response.text();
      console.log("Response:", result);
      toast.success("File uploaded and processed successfully!");
      alert(result);
    } catch (error) {
      console.error("Error uploading file:", error);
      toast.error("Failed to upload file. Please try again.");
      setErrorMessage("Failed to upload file. Please try again.");
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="bg-[#1e1e1e] h-[93vh] pb-[80px] flex flex-col justify-between">
      <h2 className="text-lg font-semibold text-amber-50 pt-14 ml-[40vh]">
        Upload a CSV File:
      </h2>
      <div className="flex flex-col items-center justify-center flex-grow">
        <div className="mb-6">
          <label
            htmlFor="tablename"
            className="block mb-2 text-sm font-medium text-white"
          >
            Table Name:
          </label>
          <input
            type="text"
            placeholder="Enter table name"
            value={tablename}
            onChange={(e) => setTypeName(e.target.value)}
            className="h-[45px] w-[400px] bg-[#1e1e1e] pl-[20px] border-[1px] font-sans text-sm border-[#6b6b6b]"
          />
        </div>
        <form
          onSubmit={handleSubmit}
          className="border border-dashed h-[40vh] border-[#6e6e6e] w-[100vh] rounded-lg p-4 flex flex-col items-center justify-center"
        >
          <div className="mb-4 flex flex-col items-center">
            <input
              type="file"
              accept=".csv"
              onChange={handleFileChange}
              className="block w-[100%] text-sm border border-dashed border-[#6e6e6e] rounded-lg cursor-pointer text-white focus:outline-none"
            />
            {errorMessage && (
              <p className="text-red-500 text-sm mt-2">{errorMessage}</p>
            )}
          </div>
        </form>
      </div>
      <div className="flex justify-end items-center">
        <button
          onClick={handleSubmit}
          disabled={uploading}
          className={`px-4 py-2 text-white rounded mr-[50vh] ${
            uploading
              ? "bg-gray-400 cursor-not-allowed"
              : "bg-blue-500 hover:bg-blue-600"
          }`}
        >
          {uploading ? "Uploading..." : "Send"}
        </button>
      </div>
    </div>
  );
};

export default CsvUploader;
