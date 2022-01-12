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
import Success from "./components/Success";

const defaultPreferences = {
  helperMode: true,
};

function App() {
  const [succeeded, setSucceeded] = useState(
    dbget("lastSuccess") === new Date().toDateString()
  );
  const [currentPage, setCurrentPage] = useState(
    dbget("lastSuccess") === new Date().toDateString()
      ? PAGES.SUCCESS
      : PAGES.INSTRUCTIONS
  );
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

  const handleSuccess = () => {
    setCurrentWord("");
    setCurrentPage(PAGES.SUCCESS);
    setSucceeded(true);
    dbset("lastSuccess", new Date().toDateString());
  };

  return (
    <div
      style={{ maxWidth: "600px", minHeight: "100vh" }}
      className={
        "is-flex is-flex-direction-column mx-auto" +
        (currentPage === PAGES.WORKBENCH
          ? " is-justify-content-space-between"
          : "")
      }
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
      ) : currentPage === PAGES.SUCCESS ? (
        <Success />
      ) : (
        <Workbench
          length={wordLength}
          letters={toTamilLetters(currentWord)}
          complete={succeeded}
          onVerified={handleVerified}
          onSuccess={handleSuccess}
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
