import React from "react";
import InputBoxes from "./InputBoxes";

function Workbench({ length }) {
  return (
    <div className="is-flex is-flex-direction-column is-justify-content-between">
      <InputBoxes length={length} word="நினைத்துள்ள" />
    </div>
  );
}

export default Workbench;
