import ReactDOM from "react-dom";
import { useState, useRef, useEffect, useContext } from "react";
import * as _ from "lodash";
import * as UC from "../unicode-utils";
import { UsedLetters, LetterHint } from "../components/used-letters";

import { IntlMsg } from "../messages-ta";

import { GameContext } from "../gameProvider";

export function Input() {
    const { gameState, guessWord, checkDuplicate } = useContext(GameContext);

    let [word, updateWord] = useState("");
    let [msg, updateMsg] = useState("");

    function validate(e) {
        e.preventDefault();
        if (_.trim(word, " ").unicodeLength() != gameState.word_length) {
            updateMsg(IntlMsg.msg_invalid_length);
        } else if (checkDuplicate(word)) {
            updateMsg(IntlMsg.msg_duplicate);
        } else {
            updateMsg("");
            guessWord(word);
        }
        return false;
    }

    let debouncedValidate = _.debounce(validate, 300);

    function onKeyUp(event) {
        if (event.key === "Enter") {
            debouncedValidate(event);
        }
    }

    console.log("posHint full", gameState.posHint);

    return (
        <div className="flex flex-col w-full justify-center items-center gap-1">
            <div className="text-pink-700">{msg}</div>
            <div>
                <input
                    className="rounded border-solid border-2 text-pink-500"
                    onChange={(e) => updateWord(e.target.value)}
                    onKeyUp={(e) => onKeyUp(e)}
                />
                <button
                    className="rounded bg-indigo-300 px-1 hover:bg-indigo-500"
                    onClick={(e) => debouncedValidate(e)}
                >
                    {IntlMsg.btn_try}
                </button>
            </div>
            <LetterHint word_length={gameState.word_length} word={word} letterStatus={gameState.letterHint} posHint={gameState.posHint} />
        </div>
    );
}

export function InputArea() {
    const { gameState, persistGameState, server, end_point, showSuccess } = useContext(GameContext);
    return (
        <div>
            {!gameState.over ? (
                <Input />
            ) : (
                <div className="flex mx-auto justify-center">
                    <button
                        onClick={(e) => showSuccess()}
                        className="rounded bg-indigo-600 hover:bg-indigo-200 p-1 text-white"
                    >
                        {IntlMsg.btn_game_over}
                    </button>
                </div>
            )}
        </div>
    );
}
