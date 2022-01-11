import React, { useState } from "react";
import { layout as t99layout } from "./tamil99";
import KeyRow from "./KeyRow";
import { diacritics, vowels } from "../utils";

const ROWTYPE = { MAIN: "main", ALT: "alt" };

function Keyboard({ onType }) {
  const [rowType, setRowType] = useState(ROWTYPE.MAIN);
  let rows = t99layout.rows;
  let altrows = t99layout.altrows;

  const handleType = (char) => {
    !diacritics[char] && !vowels[char] && char !== "\u2190" && char !== ""
      ? setRowType(ROWTYPE.ALT)
      : setRowType(ROWTYPE.MAIN);
    onType(char);
  };

  return (
    <div id="keyboard">
      {(rowType === ROWTYPE.MAIN ? rows : altrows).map((row, i) => (
        <KeyRow key={i} glyphs={row.split(" ")} onType={(c) => handleType(c)} />
      ))}
    </div>
  );
}

export default Keyboard;
