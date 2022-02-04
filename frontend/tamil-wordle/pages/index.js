import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

import * as _ from "lodash";
import ReactDOM from "react-dom";
import * as UC from "../unicode-utils";
import { IntlMsg } from "../messages-ta";

import { InputArea } from "../components/word-input";
import { Tile, Tiles } from "../components/tiles";

import { useState, useRef, useEffect, useContext } from "react";

import { zonedTimeToUtc } from "date-fns-tz";
import { isAfter, sub, differenceInDays, differenceInMinutes } from "date-fns";

import { GameContext, GameProvider } from "../gameProvider";

export function Game() {
    const { gameState, persistGameState, server, end_point, showSuccess } = useContext(GameContext);

    return (
        <div className="flex flex-col justify-center space-y-2">
            <div className="flex flex-col justify-center space-y-2">
                <div className="flex flex-grow justify-center">
                    <Tiles word_length={gameState.word_length} words={gameState.words} />
                </div>
            </div>
            <InputArea />
        </div>
    );
}

export default function Home({ word_length, server, end_point, loadError }) {
    return (
        <div className="">
            <Head>
                <title>Tamil Wordle</title>
                <meta name="description" content="A game with tamil words" />
                <link rel="icon" href="/favicon.ico" />
                <link rel="preload" href="/fonts/NotoSansTamil/NotoSansTamil-Regular.ttf" as="font" crossOrigin="" />
            </Head>

            <div className="container flex flex-col mx-auto h-screen">
                <main className="main grow">
                    {loadError ? (
                        <div className="rounded bg-pink-300 bold">{loadError}</div>
                    ) : (
                        <GameProvider server={server} end_point={end_point}>
                            <Game loadError={loadError} />
                        </GameProvider>
                    )}
                </main>

                <footer>
                    <hr />
                    <div className="flex flex-row space-x-2">
                        <img src="/sol.png" height="32" width="32" />
                        <div className="grow">&nbsp;</div>
                        <div>
                            <a href="https://github.com/psankar/tamil-wordle" className="underline">
                                Gidhub Project
                            </a>
                        </div>
                    </div>
                </footer>
            </div>
        </div>
    );
}

export async function getServerSideProps(context) {
    const server = process.env.backend_server;
    const end_point = process.env.end_point;
    try {
        const res = await fetch(`${server}/get-current-word-len`);
        const data = await res.json();

        if (!data) {
            return {
                notFound: true,
            };
        }

        return {
            props: {
                word_length: data.Length,
                server,
                end_point,
            },
        };
    } catch (err) {
        console.log(err);
        return {
            props: {
                loadError: "Error communicating with server",
            },
        };
    }
}
