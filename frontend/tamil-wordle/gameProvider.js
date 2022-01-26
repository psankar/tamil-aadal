import React from "react";
import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import { States } from "./game";

import { getLetterPos } from "./tamil-letters";

import { Modal } from "./components/modal";
import { Success } from "./components/success-message";
import { Alert } from "./components/alert";
import { Title } from "./components/title";
import { Help } from "./components/help-page";

import { zonedTimeToUtc } from "date-fns-tz";
import { isAfter, sub, differenceInDays, differenceInMinutes } from "date-fns";

let initialState = {
    updated: new Date(),
    showHelp: true,
    showOnStart: true,
    over: false,
    words: [], // [{word, status}]
    triedWords: {}, // map of tried Words for checking duplicates
    letterHint: {}, // {leter: [CORRECT, WRONG_PLACE, NOT THERE] for the given pos
    posHint: [], // [ [row, col] ] - for each pos, holds the row/col match in the 19x13 tamil letter matrix
};

export const GameContext = React.createContext(initialState);

export function GameProvider(props) {
    const [gameState, updateGameState] = useState(initialState);
    const [showModal, updateShowModal] = useState(false);
    const [alert, updateAlert] = useState({ msg: "", show: false, status: "error" });

    function persistGameState(state) {
        updateGameState(state);
        if (window && window.localStorage) {
            window.localStorage.setItem(props.end_point, JSON.stringify({ ...state, updated: new Date() }));
        }
    }

    async function guessWord(guess) {
        gameState.triedWords[guess] = true;
        persistGameState({ ...gameState });
        let word = [];
        guess.forUnicodeEach((x) => word.push(x));
        try {
            const res = await fetch(`${props.server}/${props.end_point}`, {
                method: "POST",
                mode: "cors",
                cache: "no-cache",
                headers: { "Content-Type": "application/json", Accept: "*/*" },
                body: JSON.stringify(word),
            });
            if (res.status === 200) {
                let data = await res.json();
                gameState.words.push({ word: guess, result: data });
                let pos = 0;
                guess.forUnicodeEach((ch) => {
                    if (!gameState.letterHint[ch]) gameState.letterHint[ch] = [];

                    // update the hint
                    let hint = gameState.letterHint[ch];
                    if (hint.length < gameState.word_length) {
                        for (let i = hint.length; i < gameState.word_length; i++) hint.push(States.LETTER_UNKNOWN);
                    }
                    if (hint[pos] === States.LETTER_UNKNOWN) hint[pos] = data[pos][0];
                    if (data[pos] === States.LETTER_NOT_FOUND) {
                        hint.fill(States.LETTER_NOT_FOUND);
                    }

                    // update pos hints
                    if (gameState.posHint.length <= i + 1) {
                        gameState.posHint.push([-1, -1]);
                        gameState.posHint.push([-1, -1]);
                    }
                    if (data[pos].length > 1) {
                        let posHint = gameState.posHint[pos];
                        if (data[pos][1] === States.MEI_MATCHED) posHint[1] = getLetterPos(ch)[1];
                        else if (data[pos][1] === States.UYIR_MATCHED) {
                            posHint[0] = getLetterPos(ch)[0];
                        }
                    }

                    pos += 1;
                });
                persistGameState({ ...gameState });
            } else if (res.status === 202) {
                let data = [];
                let i = 0;
                guess.forUnicodeEach((x) => {
                    data.push([States.LETTER_MATCHED]);
                    gameState.posHint[i] = getLetterPos[x];
                    i += 1;
                });
                if (!gameState.over) {
                    gameState.words.push({ word: guess, result: data });
                    gameState.over = true;
                    persistGameState({ ...gameState });
                }
                onGameOver();
            }
        } catch (error) {
            gameState.triedWords[guess] = undefined;
            persistGameState({ ...gameState });
            showAlert("error", error);
            console.error(error);
        }
    }

    function onGameOver() {
        updateShowModal(true);
    }

    function checkDuplicate(word) {
        return gameState.triedWords[word] !== undefined;
    }
    function showError(error) {}
    function showHelp(show) {
        updateGameState({ ...gameState, showHelp: show });
    }
    function toReset(state, end_point) {
        let startTime = { hours: 3, minutes: 30 };
        let reset = false;
        let today = new Date();
        today = sub(today, startTime);
        // reset the state everyday at 9:00 IST
        let start = new Date();
        start.setUTCHours(startTime.hours);
        start.setUTCMinutes(startTime.minutes);
        start = sub(start, startTime);

        if (differenceInMinutes(today, start) < 0) {
            start = sub(start, { hours: 24 });
        }

        if (state && state.updated) {
            let lastUpdated = new Date(state.updated);
            lastUpdated = sub(lastUpdated, startTime);
            if (!isAfter(lastUpdated, start)) {
                reset = true;
            }
        }
        return reset;
    }

    let showAlert = (status, msg) => updateAlert({ ...alert, msg: msg + "", status, show: true });

    function showSuccess() {
        updateShowModal(true);
    }

    // update game word length for the day
    useEffect(async () => {
        console.log("loading gamestate...", new Date(gameState?.updated).toUTCString());
        let state = window.localStorage.getItem(props.end_point);
        let gs = JSON.parse(state) || { ...initialState, word_length: props.word_length };
        if (toReset(gs, props.end_point)) {
            reset = true;
            console.log("resetting");
            gs = { ...initialState, word_length: props.word_length };
        }
        gameState = gs;
        updateGameState({ ...gs });

        const res = await fetch(`${props.server}/get-current-word-len`);
        const data = await res.json();
        updateGameState({ ...gameState, word_length: data.Length, showOnStart: false });
    }, []);

    return (
        <GameContext.Provider
            value={{
                gameState,
                showHelp,
                persistGameState,
                server: props.server,
                end_point: props.end_point,
                checkDuplicate,
                showSuccess,
                guessWord,
            }}
        >
            <Title />
            <Alert status={alert.status} show={alert.show} onHide={() => updateAlert({ ...alert, show: false })}>
                {alert.msg}
            </Alert>
            <div>{props.children}</div>
            <Help />
            <Modal show={showModal} onClose={() => updateShowModal(false)}>
                <Success word_length={gameState.word_length} words={gameState.words} />
            </Modal>
        </GameContext.Provider>
    );
}
