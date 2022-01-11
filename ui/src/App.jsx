import React, { useState, useEffect } from "react";
import { get as dbget, set as dbset } from "lockr";
import "./App.css";
import Header from "./components/Header";
import Workbench from "./components/Workbench";
import Instructions from "./components/Instructions";
import axios from "axios";
import Keyboard from "./components/Keyboard";
import { diacritics, toTamilLetters } from "./utils";

function App() {
  const [hideInstructions, setHideInstructions] = useState(
    dbget("hideInstructions")
  );
  const [wordLength, setWordLength] = useState(5);
  const [lengthLoaded, setLengthLoaded] = useState(false);
  const [currentWord, setCurrentWord] = useState("");

  useEffect(() => {
    if (!lengthLoaded) {
      axios
        .get("https://tamilwordle-maleycpqdq-el.a.run.app/get-current-word-len")
        .then((res) => {
          setWordLength(res.data.Length);
          setLengthLoaded(true);
        })
        .catch((e) => {
          console.log(e);
          alert("Failed to load app");
        });
    }
  }, [lengthLoaded]);

  useEffect(() => {
    dbset("hideInstructions", hideInstructions);
  }, [hideInstructions]);

  const typeChar = (c) => {
    let letters = toTamilLetters(currentWord);

    if (c === "\u2190") {
      setCurrentWord(currentWord.slice(0, currentWord.length - 1));
    } else if (
      letters.length < wordLength ||
      // last letter + a diacritic
      (letters.length === wordLength &&
        letters[letters.length - 1].length === 1 &&
        diacritics[c])
    ) {
      setCurrentWord(currentWord + c);
    }
  };

  return (
    <div
      style={{ maxWidth: "600px" }}
      className="is-flex is-flex-direction-column mx-auto"
    >
      <Header onShowInstructions={() => setHideInstructions(false)} />
      {!lengthLoaded ? <section className="section">Loading...</section> : null}
      {hideInstructions ? (
        <Workbench
          length={wordLength}
          letters={toTamilLetters(currentWord)}
          onVerified={() => setCurrentWord("")}
        />
      ) : (
        <Instructions onHide={() => setHideInstructions(true)} />
      )}

      <Keyboard onType={(c) => typeChar(c)} />
    </div>
  );
}

export default App;
