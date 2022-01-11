import React, { useState, useEffect } from "react";
import InputBoxes from "./InputBoxes";
import axios from "axios";
import HistoryBoxes from "./HistoryBoxes";

function Workbench({ length, letters, onVerified }) {
  const [guesses, setGuesses] = useState([]);
  const [highlightEmpty, setHighlightEmpty] = useState(false);

  const verify = () => {
    if (letters.length !== length) {
      setHighlightEmpty(true);
      return;
    }
    // call the API only when all the boxes are full
    axios
      .post("https://tamilwordle-maleycpqdq-el.a.run.app/verify-word", letters)
      .then((res) => {
        setGuesses([...guesses, { letters, results: res.data }]);
      })
      .catch((e) => {
        console.log(e);
        alert("Error");
      })
      .finally(() => onVerified());
  };

  // reset red border when user starts to type again
  useEffect(() => {
    setHighlightEmpty(false);
  }, [letters]);

  return (
    <div className="is-flex is-flex-direction-column is-justify-content-between workbench">
      <div id="historyboxes">
        {guesses.map((g, i) => (
          <HistoryBoxes key={i} guess={g} />
        ))}
      </div>
      <div>
        <InputBoxes
          length={length}
          letters={letters}
          highlightEmpty={highlightEmpty}
        />
        <div className="my-3 buttons">
          <button
            className="button is-primary mx-auto"
            onClick={() => verify()}
          >
            சரிபார்
          </button>
        </div>
      </div>
    </div>
  );
}

export default Workbench;
