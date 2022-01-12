import React, { useState } from "react";
import { FiArrowLeft } from "react-icons/fi";
import Cell from "./Cell";
import { FiArrowRight } from "react-icons/fi";

function Settings({ settings, onUpdate, onClose }) {
  const [helperMode, setHelperMode] = useState(settings.helperMode);
  const [disableKeys, setDisableKeys] = useState(settings.disableKeys);

  const helperModeChanged = (e) => {
    setHelperMode(e.target.checked);
    onUpdate("helperMode", e.target.checked);
  };

  const disableKeysChanged = (e) => {
    setDisableKeys(e.target.checked);
    onUpdate("disableKeys", e.target.checked);
  };

  return (
    <section className="section">
      <h5 className="is-size-6 has-text-weight-bold">
        முந்தய முயற்சியில்த் தவறெனத் தெரிந்த எழுத்துக்களை
      </h5>
      <div className="my-4">
        <label className="checkbox" style={{ lineHeight: 1.5 }}>
          <input
            type="checkbox"
            checked={helperMode}
            onChange={(e) => helperModeChanged(e)}
          />
          &nbsp;உள்ளிடும் பொழுதுச் சுட்டிக்காட்டு
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

      <div className="my-4">
        <label className="checkbox" style={{ lineHeight: 1.5 }}>
          <input
            type="checkbox"
            checked={disableKeys}
            onChange={(e) => disableKeysChanged(e)}
          />
          &nbsp;தட்டச்சுப் பலகையில் செயலிழக்கச் செய்துவிடு
        </label>
        <div className="is-flex is-justify-content-center mt-2">
          <button className="key" style={{ maxWidth: "2rem" }}>
            அ
          </button>
          <div className="p-2">
            <span className="icon pt-3">
              <FiArrowRight />
            </span>
          </div>
          <button className="key" style={{ maxWidth: "2rem" }} disabled={true}>
            அ
          </button>
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
