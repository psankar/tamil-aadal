import * as _ from "lodash";
import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import * as UC from "../unicode-utils";

export function Tile({ letter, color, isResult = false }) {
    let st = `tile-${color}`;
    return <div className={st}>{isResult ? String.fromCodePoint(0x1f7e9) : letter}</div>;
}

export function Tiles({ words, word_length, isResult = false, heading = true }) {
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
            let color = "notthere";
            let emoji = String.fromCodePoint(0x2b1b);
            if (result[i] === "LETTER_ELSEWHERE") {
                color = "jumbled";
                emoji = String.fromCodePoint(0x1f7e8);
            } else if (result[i] === "LETTER_MATCHED") {
                color = "correct";
                emoji = String.fromCodePoint(0x1f7e9);
            }
            if (!isResult) {
                wordTiles.push(<Tile key={`key-${w}-${i}`} letter={w} color={color}></Tile>);
            } else {
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
        shareText = `Tamil Wordle (${words.length} tries)\n${tileMatrix}`;
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
        <div ref={resultTilesPreRef} className="space-x-2">
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
