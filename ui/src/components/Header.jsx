import React from "react";
import { FiHelpCircle, FiSettings } from "react-icons/fi";

function Header({ onShowInstructions }) {
  return (
    <div
      id="header"
      className="is-flex is-justify-content-space-between"
      style={{ boxShadow: "0px 1px 2px #dcdcdc" }}
    >
      <div className="py-2">
        <button
          className="button is-white has-text-grey-light"
          onClick={() => onShowInstructions()}
        >
          <span className="icon">
            <FiHelpCircle />
          </span>
        </button>
      </div>
      <div className="my-3 px-3">
        <h1 className="is-size-6 has-text-weight-bold">தமிழ் வோர்டில்</h1>
      </div>
      <div className="py-2">
        <button className="button is-white has-text-grey-light">
          <span className="icon">
            <FiSettings />
          </span>
        </button>
      </div>
    </div>
  );
}

export default Header;
