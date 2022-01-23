import { useState, useRef, useEffect } from "react";
export default function (req, res) {
    useEffect(() => {
        if(window && window.localStorage) {
            window.localStorage.removeItem("verify-word-with-uyirmei");
            window.localStorage.removeItem("verify-word");
            setTimeout(() => {
                window.location.href = "/";
            }, 3000);
        }
    });
    return <div>Resetting storage...</div>
}
