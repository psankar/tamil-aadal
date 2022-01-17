import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import * as UC from "../unicode-utils";
import { Tile, Tiles } from "../components/tiles";
import { IntlMsg } from "../messages-ta";
import { States } from "../game";

export function Help({ show, onClose, word_length=3 }) {
    return show ? (
        <div className="overlay">
            <div className="flex flex-col px-1 space-y-3 auto w-80 h-100 flex-grow">
                <div className="space-y-2">
                    <h2 className="text-2xl">விளையாடும் முறை</h2>
                    <div>மறைந்திருக்கும் சொல்லை கண்டு பிடிக்கவும்! </div>
                    <div>{word_length} எழுத்து(க்)கள் அளவு நீளமான தமிழ்ச்சொல்லை, தமிழில் தட்டச்சு செய்து, &apos;சரி பார்க்க&apos; பொத்தானை அழுத்தவும்</div>
                    <div>ஒவ்வொரு முறையும், நீங்கள் முயற்சித்த சொல்லின் எழுத்துக்கள் கொண்ட கட்டத்தின் வண்ணம் கீழ்கண்டவாறு மாறும்</div>
                    <hr/>
                    எடுத்துக்காட்டாக...
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "வானம்", result: [[States.LETTER_NOT_FOUND], [States.LETTER_NOT_FOUND], [States.LETTER_MATCHED]]}]} heading={false}/>
                    </div>
                    &apos;ம்&apos; எழுத்து சரியான இடத்தில் உள்ளது. மற்ற இரு எழுத்துக்கள் தவறு
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "சாதம்", result: [[States.LETTER_NOT_FOUND], [States.LETTER_ELSEWHERE], [States.LETTER_MATCHED]]}]} heading={false}/>
                    </div>
                    &apos;த&apos; எழுத்து சரி, ஆனால் இடம்மாறி உள்ளது.
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "தசம்", result: [[States.LETTER_MATCHED], [States.LETTER_MATCHED], [States.LETTER_MATCHED]]}]} heading={false}/>
                    </div>
                    எல்லா எழுத்துக்களும், சொல்லும் சரி. நீங்கள் வெற்றி பெற்றுவிட்டீர்கள்!
                    <br/>
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "சாதம்", result: [[States.LETTER_NOT_FOUND, States.UYIR_MATCHED], [States.LETTER_MATCHED], [States.LETTER_MATCHED]]}]} heading={false}/>
                    </div>
                    பச்சை கீழ் நோக்கிய பட்டிகள், சா=ச்+ஆ எனப் பிரித்து, &apos;ஆ&apos; உயிர்ப்பொருத்தம் உள்ளது. ஆ வரிசையில் கா, நா, தா என கீழ் நோக்கிய வரிசையில் வேறு மெய்யுடன் முயற்சி செய்யவும்.  இது எந்த இடத்தில் வருகிறதோ அந்த &apos;இடத்திற்கு&apos; மட்டுமே, எழுத்துக்கு அல்ல.
                    <div className="flex" >
                        <Tiles word_length="3" words={[{word: "சாதம்", result: [[States.LETTER_NOT_FOUND, States.MEI_MATCHED], [States.LETTER_MATCHED], [States.LETTER_MATCHED]]}]} heading={false}/>
                    </div>
                    பச்சை இடம் வலம் நோக்கிய பட்டிகள், சா=ச்+ஆ எனப் பிரித்து, &apos;ச&apos; மெய் பொருத்தம் உள்ளது. சா, சி, சு வரிசையில் வேறு மெய்யுடன் முயற்சி செய்யவும்.  இது எந்த இடத்தில் வருகிறதோ அந்த இடத்திற்கு மட்டுமே, எழுத்துக்கு அல்ல.
                    
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
