import React from "react";
import { get as dbget } from "lockr";

const ICONS = {
  LETTER_NOT_FOUND: "⚫",
  LETTER_ELSEWHERE: "🟡",
  LETTER_MATCHED: "🟢",
};

function Success() {
  const historykey = new Date().toDateString().replace(/ /g, "-");
  let guesses = dbget("guessHistory")[historykey];

  const triggerShare = () => {
    let d = new Date();
    let history = guesses
      .map(({ results }) =>
        results.reduce((acc, curr) => acc + ICONS[curr], "")
      )
      .join("\n");
    let text =
      `தமிழ் வோர்டில் (${d.getDate()}/${d.getMonth() + 1})\n\n` + history;
    if (navigator.share) {
      navigator.share({ text });
    } else {
      navigator.clipboard.writeText(text);
      alert("Content copied to clipboard");
    }
  };

  return (
    <div className="card">
      <header className="card-header">
        <p className="card-header-title">
          <span className="has-text-centered mx-auto">
            🎉️ வாழ்த்துகள்! சரியான சொல்லைக் கண்டுபிடித்துவிட்டீர்கள் 🎊️
          </span>
        </p>
      </header>
      <div className="card-content">
        <p className="has-text-centered">உங்களின் முயற்சி வரலாறு</p>
        <div className="is-flex is-flex-direction-column">
          {guesses.map(({ results }, i) => (
            <div key={i}>
              {results.reduce((acc, curr) => acc + ICONS[curr], "")}
            </div>
          ))}
        </div>
      </div>
      <footer className="card-footer">
        <button
          href="#"
          className="card-footer-item is-ghost"
          onClick={() => triggerShare()}
        >
          📣️ பகிர்
        </button>
      </footer>
    </div>
  );
}

export default Success;
