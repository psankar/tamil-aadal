import React from "react";
import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";

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

    function persistGameState(state) {
        updateGameState(state);
        if (window && window.localStorage) {
            window.localStorage.setItem(props.end_point, JSON.stringify({ ...state, updated: new Date() }));
        }
    }

    async function guessWord(guess) {}
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
            value={{ gameState, showHelp, persistGameState, server: props.server, end_point: props.end_point, checkDuplicate }}
        >
            <div>{props.children}</div>
        </GameContext.Provider>
    );
}
