import React from "react";
import { toTamilLetters } from "../utils";
import Cell from "./Cell";

function InputBoxes({ word, length }) {
  let _l = toTamilLetters(word);
  let letters = [...Array(length)].map((_, i) => _l[i] || "");
  return (
    <div className="is-flex is-justify-content-center">
      {letters.map((l, i) => (
        <Cell key={i} letter={l} />
      ))}
    </div>
  );
}

export default InputBoxes;
