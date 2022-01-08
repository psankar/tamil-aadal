import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import * as UC from "../unicode-utils";
import { Tile, Tiles } from "../components/tiles";
import { IntlMsg } from "../messages-ta";

export function Help({ show, onClose, word_length=3 }) {
    return show ? (
        <div className="overlay">
            <div className="flex flex-col px-1 space-y-3 m-auto w-80">
                <div className="space-y-2">
                    <h2 className="text-2xl">விளையாடும் முறை</h2>
                    <div>மறைந்திருக்கும் சொல்லை கண்டு பிடிக்கவும்! </div>
                    <div>{word_length} எழுத்து(க்)கள் அளவு நீளமான தமிழ்ச்சொல்லை, தமிழில் தட்டச்சு செய்து, &apos;சரி பார்க்க&apos; பொத்தானை அழுத்தவும்</div>
                    <div>ஒவ்வொரு முறையும், நீங்கள் முயற்சித்த சொல்லின் எழுத்துக்கள் கொண்ட கட்டத்தின் வண்ணம் கீழ்கண்டவாறு மாறும்</div>
                    <hr/>
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "வானம்", result: ["LETTER_UNMATCHED", "LETTER_UNMATCHED", "LETTER_MATCHED"]}]} heading={false}/>
                    </div>
                    &apos;ம்&apos; எழுத்து சரியான இடத்தில் உள்ளது. மற்ற இரு எழுத்துக்கள் தவறு
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "சாதம்", result: ["LETTER_UNMATCHED", "LETTER_ELSEWHERE", "LETTER_MATCHED"]}]} heading={false}/>
                    </div>
                    &apos;த&apos; எழுத்து சரி, ஆனால் இடம்மாறி உள்ளது.
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "தசம்", result: ["LETTER_MATCHED", "LETTER_MATCHED", "LETTER_MATCHED"]}]} heading={false}/>
                    </div>
                    எல்லா எழுத்துக்களும், சொல்லும் சரி. நீங்கள் வெற்றி பெற்றுவிட்டீர்கள்!
                </div>
                <div className="self-center">
                    <button className="rounded bg-green-300 text-blue-800 px-1" onClick={(e) => onClose()}>
                    Close
                    </button>
                </div>
            </div>
        </div>
    ) : null;
}
