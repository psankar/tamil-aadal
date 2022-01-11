import React from "react";

function KeyRow({ glyphs, onType }) {
  return (
    <div className="my-1 is-flex is-justify-content-center">
      {glyphs.map((g) => (
        <button className={`key ${g}`} key={g} onClick={() => onType(g)}>
          {g}
        </button>
      ))}
    </div>
  );
}

export default KeyRow;
