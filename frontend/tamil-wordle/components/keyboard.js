import ReactDOM from "react-dom";
import { useState, useRef, useEffect, useContext } from "react";
import * as _ from "lodash";
import { Layout } from "../layout";

import { Letters, getLetterPos } from "../tamil-letters";

export function KeyRow({ row, onKeyClick }) {
    let keys = [];
    let items = row.split(" ");
    items.forEach((letter) => {
        keys.push(<Key value={letter} onClick={onKeyClick} />);
    });
    return <div className="ps-1 space-x-1 flex justify-items-stretch justify-center">{keys}</div>;
}

export function Key({ value, onClick }) {
    return (
        <button
            className="border rounded px-1 justify-center justify-items-center hover:bg-gray-300"
            onClick={(e) => onClick(value)}
        >
            {value}
        </button>
    );
}

export function useKeyboard(text, updateText) {
    function onButtonClick(keyClicked) {
        let toUpdate = text + keyClicked;
        if(keyClicked === "clear") {
            toUpdate = "";
        } else if (keyClicked === "space") {
            toUpdate = text + " ";
        } else if (keyClicked === "backspace") {
            if (text.length > 0) {
                let last = text[text.length - 1];
                let pos = getLetterPos(last);
                if (last === "\u0bcd") {
                    toUpdate = text.slice(0, text.length - 2);
                } else if (pos && pos[0] != 0) {
                    // non uyir
                    if (pos[1] == 12) {
                        // mei
                        toUpdate = text.slice(0, text.length - 1);
                    } else if (pos[1] > 0) {
                        // uyir mei variation
                        toUpdate = text.slice(0, text.length - 1) + Letters[pos[0]][0];
                    } else {
                        toUpdate = text.slice(0, text.length - 1) + Letters[pos[0]][12];
                    }
                } else toUpdate = text.slice(0, text.length - 1);
            } else toUpdate = text.slice(0, text.length - 1);
        } else if (text.length > 0) {
            let last = text[text.length - 1];
            let pos = getLetterPos(last);
            let pressedPos = getLetterPos(keyClicked);
            if (pos && pressedPos) {
                if (pos[0] > 0 && pressedPos[0] == 0) {
                    toUpdate = text.slice(0, text.length - 1) + Letters[pos[0]][pressedPos[1]];
                }
            }
        }
        updateText(toUpdate);
    }

    function OnScreenKeyboard() {
        let keys = [];
        Layout.normal.forEach((row) => {
            keys.push(<KeyRow row={row} onKeyClick={onButtonClick} />);
        });
        return (
            <div className="p-1 flex flex-col justify-center">
                <div className="flex grow flex-col justify-center justify-items-center ">{keys}</div>
            </div>
        );
    }

    return {OnScreenKeyboard}
}
