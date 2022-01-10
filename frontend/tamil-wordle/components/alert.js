export function Alert({ status, msg, show, onHide }) {
    let cl = `alert-${status}`
    return show ? (
        <div className={cl} role="alert">
            {msg}
            <button
                type="button"
                className="btn-close box-content w-4 h-4 p-1 ml-auto text-yellow-900 border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-yellow-900 hover:opacity-75 hover:no-underline"
                data-bs-dismiss="alert"
                aria-label="Close"
                onClick={(e) => onHide()}
            >
                x
            </button>
        </div>
    ) : null;
}
