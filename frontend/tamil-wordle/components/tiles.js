import * as _ from "lodash";
import ReactDOM from "react-dom";
import { useState, useRef, useEffect, useContext } from "react";
import * as UC from "../unicode-utils";
import { States } from "../game";
import { getLetterPos } from "../tamil-letters";

import { IntlMsg } from "../messages-ta";

import { GameContext } from "../gameProvider";

const emojimap = {
    LETTER_NOT_FOUND: { UYIR_MATCHED: 0x1f5a4, MEI_MATCHED: 0x26ab },
    LETTER_MATCHED: { UYIR_MATCHED: 0x1f49a, MEI_MATCHED: 0x1f7e2 },
    LETTER_ELSEWHERE: { UYIR_MATCHED: 0x1f49b, MEI_MATCHED: 0x1f7e1 },
    LETTER_UNKNOWN: { UYIR_MATCHED: 0x1f49c, MEI_MATCHED: 0x1f7e3 },
};

function mapStateToUIProperties(letterState, posState, uyirMeiHintsUsed = false) {
    let color = "notthere";
    let anim = "animate-flip";
    let emoji = String.fromCodePoint(0x2b1b);
    let border = "";
    if (letterState === States.LETTER_ELSEWHERE) {
        color = "jumbled";
        anim = "animate-bounce";
        emoji = String.fromCodePoint(0x1f7e8);
    } else if (letterState === States.LETTER_MATCHED) {
        color = "correct";
        anim = "animate-none";
        emoji = String.fromCodePoint(0x1f7e9);
    } else if (letterState === States.LETTER_NOT_FOUND) {
        color = "notthere";
        anim = "animate-focus";
        emoji = String.fromCodePoint(0x2b1b);
    } else if (letterState === States.LETTER_UNKNOWN) {
        color = "unknown";
        anim = "animate-flip";
        emoji = String.fromCodePoint(0x2b1b);
    }
    if (posState === States.UYIR_MATCHED) {
        border = "border-x-4 border-green-500";
    } else if (posState === States.MEI_MATCHED) {
        border = "border-y-4 border-green-500";
    }

    if (uyirMeiHintsUsed && posState) {
        let cp = emojimap[letterState][posState];
        if (cp && !isNaN(cp)) emoji = String.fromCodePoint(cp);
    }

    return { color, anim, emoji, border };
}

export function Tile({
    letter,
    letterState,
    posState,
    globalLetterState,
    isResult = false,
    anim = "animate-none",
    isHint = false,
    forHelpPage = false,
}) {
    const { gameState } = useContext(GameContext);

    let order = { unknown: 0, notthere: 1, jumbled: 2, correct: 3 };
    let { color, border, emoji } = mapStateToUIProperties(letterState, posState, gameState.showUyirMeiHints);
    if (globalLetterState && globalLetterState[letter]) {
        globalLetterState[letter].forEach((st) => {
            let m = mapStateToUIProperties(st);
            if (order[color] < order[m.color]) {
                color = m.color;
            }
        });
    }
    console.log("posHint", letter, getLetterPos(letter), posState);
    if (isHint && posState) {
        let letterPos = getLetterPos(letter);
        if (letterPos && posState[0] === letterPos[0]) {
            border = "border-y-4 border-green-500";
        }
        if (letterPos && posState[1] === letterPos[1]) {
            border = "border-x-4 border-green-500";
        }
    }
    let st = `tile-${color} ${anim}`;
    if (gameState.showUyirMeiHints || forHelpPage) {
        st += ` ${border}`;
    }
    return <div className={st}>{isResult ? String.fromCodePoint(0x1f7e9) : letter}</div>;
}

export function Tiles({ words, word_length, isResult = false, heading = true, forHelpPage = false }) {
    const { gameState } = useContext(GameContext);

    const divEl = useRef(null);
    const resultTilesPreRef = useRef(null);
    let wordTiles = [];
    if (!isResult && heading) {
        for (let i = 1; i <= word_length; i++) {
            wordTiles.push(<Tile letter={i} color="gray" />);
        }
    }
    words.forEach(({ word, result }) => {
        let i = 0;
        word.forUnicodeEach((w) => {
            if (!isResult) {
                wordTiles.push(
                    <Tile
                        key={`key-${w}-${i}`}
                        letterState={result[i][0]}
                        posState={result[i].length > 1 ? result[i][1] : undefined}
                        letter={w}
                        forHelpPage={forHelpPage}
                    ></Tile>
                );
            } else {
                let { emoji } = mapStateToUIProperties(
                    result[i][0],
                    result[i].length > 1 ? result[i][1] : undefined,
                    gameState.showUyirMeiHints
                );
                wordTiles.push(emoji);
            }
            i += 1;
        });
    });
    useEffect(() => {
        if (divEl && divEl.current) {
            divEl.current.scrollTop = divEl.current.scrollHeight;
        }
    });

    let shareText = "";
    if (isResult) {
        let tileMatrix = _.join(
            _.map(_.chunk(wordTiles, word_length), (x) => _.join(x, "")),
            "\n"
        );
        shareText = `${IntlMsg.game_name} (${words.length} tries)\n${tileMatrix}`;
    }

    async function OnCopyClick() {
        await navigator.clipboard.writeText(shareText);
        alert(`Copied to clipboard! Use your favourite tool to share!\n\n${shareText}`);
    }
    async function onShareClick() {
        if (navigator.share) {
            await navigator.share(shareText);
        } else {
            await OnCopyClick();
        }
    }

    let st = "g" + word_length;
    return !isResult ? (
        <div className={st} ref={divEl}>
            {wordTiles}
        </div>
    ) : (
        <div ref={resultTilesPreRef} className="space-x-2 space-y-2">
            <pre>{shareText}</pre>
            <button
                className="rounded bg-green-300 p-1 text-blue-800 hover:bg-green-500"
                onClick={(e) => OnCopyClick()}
            >
                Copy
            </button>
            {navigator.share ? (
                <button
                    className="rounded bg-green-300 p-1 text-blue-800 hover:bg-green-500"
                    onClick={(e) => onShareClick()}
                >
                    Share
                </button>
            ) : null}
        </div>
    );
}

export function TilesHint({ word, word_length, status, letterStatus, posHint }) {
    let hint = [];
    let i = 0;
    //console.log(word, status, letterStatus);
    word.forUnicodeEach((x) => {
        let { color, anim } = mapStateToUIProperties(status[i]);

        if (i < word_length)
            hint.push(
                <Tile
                    letter={x}
                    letterState={status[i]}
                    posState={posHint && posHint.length > 0 ? posHint[i] : posHint}
                    globalLetterState={letterStatus}
                    anim={anim}
                    isHint={true}
                />
            );
        i += 1;
    });
    let gl = `g${word_length}`;
    return <div className={gl}>{hint}</div>;
}
