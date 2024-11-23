import React from "react";
import Header from "../components/Header";

import FileUploader from "./components/uploadcsv";
import CsvUploader from "./components/uploadcsv";

const Page = () => {
  return (
    <div>
      <Header />
      <CsvUploader/>
    </div>
  );
};

export default Page;
