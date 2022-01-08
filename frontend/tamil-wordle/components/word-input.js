import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import * as UC from "../unicode-utils";

import {IntlMsg} from "../messages-ta";

export function Input({ word_length, onNewGuess, checkDuplicate }) {
    let [word, updateWord] = useState("");
    let [msg, updateMsg] = useState("");
    const validate = (e) => {
        e.preventDefault();
        if (word.unicodeLength() != word_length) {
            updateMsg(IntlMsg.msg_invalid_length);
        } else if (checkDuplicate(word)) {
            updateMsg(IntlMsg.msg_duplicate);
        } else {
            updateMsg("");
            onNewGuess(word);
        }
        return false;
    };
    return (
        <div className="flex flex-col w-full justify-center items-center gap-1">
            <div className="text-pink-700">{msg}</div>
            <div>
                <input
                    className="rounded border-solid border-2 text-pink-500"
                    onChange={(e) => updateWord(e.target.value)}
                />
                <button className="rounded bg-indigo-300 px-1 hover:bg-indigo-500" onClick={(e) => validate(e)}>
                    {IntlMsg.btn_try}
                </button>
            </div>
        </div>
    );
}
