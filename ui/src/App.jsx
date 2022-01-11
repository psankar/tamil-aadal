import React, { useState, useEffect } from "react";
import { set as dbset, get as dbget } from "lockr";
import "./App.css";
import Header from "./components/Header";
import Workbench from "./components/Workbench";
import Instructions from "./components/Instructions";
import axios from "axios";
import Keyboard from "./components/Keyboard";
import { diacritics, toTamilLetters } from "./utils";
import Settings from "./components/Settings";
import { PAGES } from "./utils";

const defaultPreferences = {
  helperMode: true,
};

function App() {
  const [currentPage, setCurrentPage] = useState(PAGES.INSTRUCTIONS);
  const [wordLength, setWordLength] = useState(5);
  const [lengthLoaded, setLengthLoaded] = useState(false);
  const [currentWord, setCurrentWord] = useState("");
  const [blacklist, setBlackList] = useState(new Set());
  const [settings, setSettings] = useState(dbget("userPreferences"));
  if (!settings) {
    setSettings(defaultPreferences);
    dbset("userPreferences", defaultPreferences);
  }

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

  const handleVerified = ({ wrongLetters }) => {
    setCurrentWord("");
    let _black = new Set(Array.from(blacklist).concat(wrongLetters));
    setBlackList(_black);
  };

  const onUpdateSettings = (key, value) => {
    setSettings({ ...settings, [key]: value });
  };

  useEffect(() => dbset("userPreferences", settings), [settings]);

  return (
    <div
      style={{ maxWidth: "600px" }}
      className="is-flex is-flex-direction-column mx-auto"
    >
      <Header onShow={(page) => setCurrentPage(page)} />
      {!lengthLoaded ? <section className="section">Loading...</section> : null}
      {currentPage === PAGES.INSTRUCTIONS ? (
        <Instructions onHide={() => setCurrentPage(PAGES.WORKBENCH)} />
      ) : currentPage === PAGES.SETTINGS ? (
        <Settings
          settings={settings}
          onClose={() => setCurrentPage(PAGES.WORKBENCH)}
          onUpdate={onUpdateSettings}
        />
      ) : (
        <Workbench
          length={wordLength}
          letters={toTamilLetters(currentWord)}
          onVerified={handleVerified}
          blacklist={settings.helperMode ? blacklist : new Set()}
        />
      )}
      {currentPage === PAGES.WORKBENCH ? (
        <Keyboard onType={(c) => typeChar(c)} blacklist={blacklist} />
      ) : null}
    </div>
  );
}

export default App;
