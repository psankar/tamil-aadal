import { IntlMsg } from "../messages-ta";
import { GameContext } from "../gameProvider";
import { useState, useRef, useEffect, useContext } from "react";

function Questionmark() {
    return (
        <div>
            <img src="/help.png" />
            help
        </div>
    );
}
export function Title() {
    const { showHelp } = useContext(GameContext);
    return (
        <div className="self-center flex space-x-5 justify-center">
            <div className="flex flex-col justify-center">
                <h1 className="self-center text-2xl">{IntlMsg.game_name}</h1>
                <h1 className="self-center text-2xl">Tamil Wordle</h1>
            </div>
            <div>
                <a href="#" onClick={(e) => showHelp(true)}>
                    <Questionmark />
                </a>
            </div>
        </div>
    );
}
