import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";

import * as _ from "lodash";
import ReactDOM from "react-dom";
import * as UC from "../unicode-utils";
import { IntlMsg } from "../messages-ta";

import { Modal } from "../components/modal";
import { Success } from "../components/success-message";
import { Input } from "../components/word-input";
import { Tile, Tiles } from "../components/tiles";
import { UsedLetters } from "../components/used-letters";

import { useState, useRef, useEffect } from "react";

// config
const server = process.env.backend_server;

export default function Home({ word_length }) {
    let [gameState, updateGameState] = useState({
        over: false,
        word_length,
        words: [], // [{word, status}]
    }); // {word, result}
    let [showModal, updateShowModal] = useState(false);

    if (!gameState.word_length) {
        return <div>Error communicating with server</div>;
    }

    function checkDuplicate(word) {
        return _.find(gameState.words, (x) => x.word === word) !== undefined;
    }

    function onGameOver() {
        updateShowModal(true);
    }

    async function onNewGuess(guess) {
        let word = [];
        guess.forUnicodeEach((x) => word.push(x));
        try {
            const res = await fetch(`${server}/verify-word`, {
                method: "POST",
                mode: "cors",
                cache: "no-cache",
                headers: { "Content-Type": "application/json", Accept: "*/*" },
                body: JSON.stringify(word),
            });
            if (res.status === 200) {
                let data = await res.json();
                gameState.words.push({ word: guess, result: data });
                updateGameState({ ...gameState });
            } else if (res.status === 202) {
                let data = [];
                guess.forUnicodeEach((x) => {
                    data.push("LETTER_MATCHED");
                });
                if (!gameState.over) {
                    gameState.words.push({ word: guess, result: data });
                    gameState.over = true;
                    updateGameState({ ...gameState });
                }
                onGameOver();
            }
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div className="container mx-auto">
            <Head>
                <title>Tamil Wordle</title>
                <meta name="description" content="A game with tamil words" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className="main">
                <h3 className={styles.title}> {IntlMsg.game_name} </h3>
                <div className="flex flex-col justify-center space-y-2">
                    <div className="flex flex-grow justify-center">
                        <Tiles word_length={gameState.word_length} words={gameState.words} />
                    </div>
                    {!gameState.over ? (
                        <Input
                            word_length={gameState.word_length}
                            onNewGuess={onNewGuess}
                            checkDuplicate={checkDuplicate}
                            onGameOver
                        />
                    ) : (
                        <div className="flex mx-auto justify-center">
                            <button
                                onClick={(e) => updateShowModal(true)}
                                className="rounded bg-indigo-600 hover:bg-indigo-200 p-1 text-white"
                            >
                                {IntlMsg.btn_game_over}
                            </button>
                        </div>
                    )}
                    <UsedLetters />
                </div>
                <Modal show={showModal} onClose={() => updateShowModal(false)}>
                    <Success word_length={gameState.word_length} words={gameState.words} />
                </Modal>
            </main>

            <footer></footer>
        </div>
    );
}

export async function getServerSideProps(context) {
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
        },
    };
}
