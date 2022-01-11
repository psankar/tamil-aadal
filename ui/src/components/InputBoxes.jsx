import React from "react";
import Cell from "./Cell";

function InputBoxes({ letters, length, highlightEmpty }) {
  let fullarray = [...Array(length)].map((_, i) => letters[i] || "");
  return (
    <div className="is-flex is-justify-content-center">
      {fullarray.map((l, i) => (
        <Cell
          key={i}
          letter={l}
          borderColor={
            highlightEmpty && !l.length
              ? "hsl(348, 100%, 61%)"
              : l.length
              ? "hsl(0, 0%, 29%)"
              : "hsl(0, 0%, 86%)"
          }
        />
      ))}
    </div>
  );
}

export default InputBoxes;
