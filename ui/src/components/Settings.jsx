import React, { useState } from "react";
import { FiArrowLeft } from "react-icons/fi";
import Cell from "./Cell";
import { FiArrowRight } from "react-icons/fi";

function Settings({ settings, onUpdate, onClose }) {
  const [helperMode, setHelperMode] = useState(settings.helperMode);

  const helperModeChanged = (e) => {
    setHelperMode(e.target.checked);
    onUpdate("helperMode", e.target.checked);
  };

  return (
    <section className="section">
      <div className="my-3">
        <label className="checkbox" style={{ lineHeight: 1.5 }}>
          <input
            type="checkbox"
            checked={helperMode}
            onChange={(e) => helperModeChanged(e)}
          />
          &nbsp;முந்தய முயற்சியில்த் தவறெனத் தெரிந்த எழுத்துக்களைச்
          சுட்டிக்காட்டு
        </label>
        <div className="is-flex is-justify-content-center">
          <Cell letter="கெ" />
          <div className="p-4">
            <span className="icon pt-3">
              <FiArrowRight />
            </span>
          </div>
          <Cell letter="கெ" borderColor="hsl(48, 100%, 29%)" />
        </div>
      </div>

      <div className="buttons pt-4">
        <button
          className="button mx-auto is-primary has-text-weight-bold"
          onClick={() => onClose()}
        >
          <FiArrowLeft />
          &nbsp;&nbsp;விளையாட்டு
        </button>
      </div>
    </section>
  );
}

export default Settings;
