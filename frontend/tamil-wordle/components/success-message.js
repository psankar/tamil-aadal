import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import * as UC from "../unicode-utils";
import { Tile, Tiles } from "../components/tiles";
import { IntlMsg } from "../messages-ta";

export function Success({ word_length, words }) {
    let successMsg = IntlMsg.btn_final_message;
    return (
        <div>
            <div>{successMsg} </div>
            <div>
                <Tiles word_length={word_length} words={words} isResult={true} />
            </div>
        </div>
    );
}
