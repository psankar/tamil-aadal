import React, { useState, useEffect } from "react";
import { get as dbget, set as dbset } from "lockr";
import "./App.css";
import Header from "./components/Header";
import Workbench from "./components/Workbench";
import Instructions from "./components/Instructions";
import axios from "axios";
import Keyboard from "./components/Keyboard";

function App() {
  const [hideInstructions, setHideInstructions] = useState(
    dbget("hideInstructions")
  );
  const [wordLength, setWordLength] = useState(5);
  const [lengthLoaded, setLengthLoaded] = useState(false);

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

  return (
    <div
      style={{ maxWidth: "600px" }}
      className="is-flex is-flex-direction-column mx-auto"
    >
      <Header onShowInstructions={() => setHideInstructions(false)} />
      {!lengthLoaded ? <section className="section">Loading...</section> : null}
      {hideInstructions ? (
        <Workbench length={wordLength} />
      ) : (
        <Instructions onHide={() => setHideInstructions(true)} />
      )}

      <Keyboard onType={(c) => console.log(c)} />
    </div>
  );
}

export default App;
