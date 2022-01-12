import React from "react";
import Cell from "./Cell";

function InputBoxes({ letters, length, blacklist, highlightEmpty }) {
  let fullarray = [...Array(length)].map((_, i) => letters[i] || "");
  return (
    <div id="input-boxes" className="is-flex is-justify-content-center">
      {fullarray.map((l, i) => (
        <Cell
          key={i}
          letter={l}
          borderColor={
            blacklist.has(l)
              ? "hsl(48, 100%, 29%)"
              : highlightEmpty && !l.length
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
