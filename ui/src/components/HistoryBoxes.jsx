import React from "react";
import Cell from "./Cell";

function HistoryBoxes({ guess }) {
  let { letters, results } = guess;
  return (
    <div className="is-flex is-justify-content-center">
      {letters.map((l, i) => (
        <Cell key={i} letter={l} status={results[i]} />
      ))}
    </div>
  );
}

export default HistoryBoxes;
