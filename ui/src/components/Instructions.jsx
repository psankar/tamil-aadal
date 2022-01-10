import React from "react";
import Cell from "./Cell";
import { FiArrowLeft } from "react-icons/fi";

function Instructions({ onHide }) {
  return (
    <section className="section">
      <h2 className="title is-size-6 has-text-weight-semi-bold">
        "வோர்டில்" விளையாடுவது எப்படி?
      </h2>
      <p>"வோர்டில்" நினைத்துள்ள வார்த்தையை நீங்கள் சரியா யூகிக்க வேண்டும்.</p>

      <p className="my-3">
        நீங்கள் ஒவ்வொரு முறையும் யூகம் செய்யும் பொழுதும் உங்களுக்கு உதவும்
        விதமாக பொட்டிகளின் நிறம் மாறும்.
      </p>

      <div style={{ margin: "2rem 0" }}>
        <hr />
      </div>

      <div className="my-3">
        <div className="is-flex is-justify-content-center">
          <Cell letter="தி" status="LETTER_MATCHED" />
          <Cell letter="ரு" />
          <Cell letter="ம" />
          <Cell letter="ண" />
          <Cell letter="ம்" />
        </div>
        <p>பெட்டி பச்சையாக மாறினால், எழுத்து பொருத்தமான இடத்தில் உள்ளது.</p>
      </div>

      <div className="my-3">
        <div className="is-flex is-justify-content-center">
          <Cell letter="வி" />
          <Cell letter="ளை" />
          <Cell letter="யா" status="LETTER_ELSEWHERE" />
          <Cell letter="ட்" />
          <Cell letter="டு" />
        </div>
        <p>பெட்டி மஞ்சளாக மாறினால், எழுத்து தவறான் இடத்தில் உள்ளது.</p>
      </div>

      <div className="my-3">
        <div className="is-flex is-justify-content-center">
          <Cell letter="செ" />
          <Cell letter="வ்" status="LETTER_NOT_FOUND" />
          <Cell letter="வா" />
          <Cell letter="ன" />
          <Cell letter="ம்" />
        </div>
        <p>
          பெட்டி இளங்கருப்பாக மாறினால், எழுத்து இன்றய வார்த்தையில் இடம்
          பிடிக்கவில்லை.
        </p>
      </div>

      <div className="buttons pt-4">
        <button
          className="button mx-auto is-primary is-large has-text-weight-bold"
          onClick={() => onHide()}
        >
          <FiArrowLeft />
          &nbsp;&nbsp;விளையாட்டு
        </button>
      </div>
    </section>
  );
}

export default Instructions;
