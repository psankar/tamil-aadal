import * as _ from "lodash";
import ReactDOM from "react-dom";
import {useContext} from 'react';
import * as UC from "../unicode-utils";
import { States } from "../game";

import {GameContext} from "../gameProvider";

import { Tile, Tiles, TilesHint } from "../components/tiles";

export function UsedLetters() {
    let letters = [];
    Letters.forEach((row) => {
        let thisRow = [];
        row.forEach((c) => {
            thisRow.push(
                <td>
                    <div className="tile-gray">{c}</div>
                </td>
            );
        });
        letters.push(<tr>{thisRow}</tr>);
    });
    return (
        <div className="flex">
            <table>{letters}</table>
        </div>
    );
}

export function LetterHint({word, word_length, letterStatus, posHint}) {
    const {getLetterStatusForWord} = useContext(GameContext);
    let status = getLetterStatusForWord(word);
    return (
        <div className="flex">
            <TilesHint word_length={word_length} word={word} status={status} letterStatus={letterStatus} posHint={posHint}/>
        </div>
    );
}
