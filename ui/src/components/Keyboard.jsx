import React, { useState } from "react";
import { layout as t99layout } from "./tamil99";
import KeyRow from "./KeyRow";
import { diacritics, vowels } from "../utils";

const ROWTYPE = { MAIN: "main", ALT: "alt" };

function Keyboard({ blacklist, onType }) {
  const [rowType, setRowType] = useState(ROWTYPE.MAIN);
  const [prevChar, setPrevChar] = useState("");

  let rows = t99layout.rows;
  let altrows = t99layout.altrows;
  const actionKeys = {
    "\u2190": true, // backspace
    "\u2191": true, // shift
  };

  const handleType = (char) => {
    if (char === "\u2191") {
      toggleLayout();
      return;
    }
    !diacritics[char] && !vowels[char] && !actionKeys[char] && char !== ""
      ? setRowType(ROWTYPE.ALT)
      : setRowType(ROWTYPE.MAIN);
    onType(char);
    setPrevChar(char);
  };

  const toggleLayout = () => {
    ROWTYPE.MAIN === rowType
      ? setRowType(ROWTYPE.ALT)
      : setRowType(ROWTYPE.MAIN);
  };

  return (
    <div id="keyboard">
      {(rowType === ROWTYPE.MAIN ? rows : altrows).map((row, i) => (
        <KeyRow
          key={i}
          glyphs={row.split(" ")}
          prevChar={prevChar}
          blacklist={blacklist}
          onType={(c) => handleType(c)}
        />
      ))}
    </div>
  );
}

export default Keyboard;
