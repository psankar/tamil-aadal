import * as _ from "lodash";
import ReactDOM from "react-dom";
import * as UC from "../unicode-utils";
import { States } from "../game";

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
    let status = _.times(word_length, _.constant(States.LETTER_UNKNOWN));
    let i = 0;
    word.forUnicodeEach(c => {
        let hint = letterStatus[c];
        if(hint && hint.length > i) {
            status[i] = hint[i];
        }
        i += 1
    });
    return (
        <div className="flex">
            <TilesHint word_length={word_length} word={word} status={status} letterStatus={letterStatus} posHint={posHint}/>
        </div>
    );
}
