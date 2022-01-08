import ReactDOM from "react-dom";
import { useState, useRef, useEffect } from "react";

export function Modal({ show, children, onClose }) {
    const [isBrowser, setIsBrowser] = useState(false);

    useEffect(() => {
        setIsBrowser(true);
    }, []);

    if (isBrowser) {
        return ReactDOM.createPortal(
            show ? (
                <div className="overlay">
                    <div className="flex flex-col px-1 space-y-3 m-auto">
                        <div>{children}</div>
                        <button className="rounded bg-green-300 text-blue-800" onClick={(e) => onClose()}>
                            Close
                        </button>
                    </div>
                </div>
            ) : null,
            document.getElementById("modal-root")
        );
    } else {
        return null;
    }
}
