import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";
import * as UC from "../unicode-utils";
import { Tile, Tiles } from "../components/tiles";

export function Success({ word_length, words }) {
    let successMsg =
        "\u0b9a\u0bb0\u0bbf\u0baf\u0bbe\u0ba9\u0020\u0b9a\u0bca\u0bb2\u0bcd\u0bb2\u0bc8\u0b95\u0bcd\u0020\u0b95\u0ba3\u0bcd\u0b9f\u0bc1\u0baa\u0bbf\u0b9f\u0bbf\u0ba4\u0bcd\u0ba4\u0bc1\u0bb5\u0bbf\u0b9f\u0bcd\u0b9f\u0bc0\u0bb0\u0bcd\u0b95\u0bb3\u0bcd\u0021\u0021\u0021 If you are interested, copy and paste the emoji table to social media";
    return (
        <div>
            <div>{successMsg} </div>
            <div>
                <Tiles word_length={word_length} words={words} isResult={true} />
            </div>
        </div>
    );
}
