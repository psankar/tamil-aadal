import Cell from "./Cell";

const steps = [
  {
    content: (
      <div>
        <h2 className="has-text-weight-bold mb-3">
          இந்த விளையாட்டு விளையாடுவது எப்படி?
        </h2>
        <p>தினம் ஒரு புது சொல்லை நீங்கள் கண்டுபிடிக்க வேண்டும்</p>
      </div>
    ),
    placement: "center",
    target: "body",
  },
  {
    content: (
      <h3>
        அந்த சொல்லில் உள்ள எழுத்துக்களுக்கு சம்மான பொட்டிகள் கொடுக்கப்பட்டு
        இருக்கும். (5 எழுத்துச்சொல் = 5 பெட்டிகள்). அத்தனை எழுத்துக்கள் உள்ள
        எதேனும் ஒரு சொல்லை யோசித்துக் கொள்ளுங்கள்.
      </h3>
    ),
    spotlightPadding: 4,
    target: "#input-boxes",
  },
  {
    content: (
      <p>
        நீங்கள் யூகித்த சொல்லைக் கீழே இருக்கும் தட்டச்சுப்பலகை மூலம் உள்ளிடவும்.
      </p>
    ),
    target: "#keyboard",
  },
  {
    content: <p>உள்ளிட்ட பின்பு, "சரிபார்" பொத்தானை அழுத்தவும்</p>,
    target: "#verify-button",
  },
  {
    content: (
      <div>
        <p>
          நீங்கள் உள்ளிட்ட சொல்லில் உள்ள எழுத்துக்கள் எப்படிப் பொருந்தியுள்ளது
          என்று பொட்டிகளின் வண்ணங்களைக் கொண்டுத் தெரிந்து கொள்ளலாம்.
        </p>
        <div className="my-3">
          <div className="is-flex is-justify-content-center">
            <Cell letter="" status="LETTER_MATCHED" />
            <p>பச்சைப்பெட்டி - எழுத்து பொருத்தமான இடத்தில் உள்ளது</p>
          </div>
        </div>
        <div className="my-3">
          <div className="is-flex is-justify-content-center">
            <Cell letter="" status="LETTER_ELSEWHERE" />
            <p>மஞ்சள் பெட்டி - எழுத்து தவறான இடத்தில் உள்ளது</p>
          </div>
        </div>

        <div className="my-3">
          <div className="is-flex is-justify-content-center">
            <Cell letter="" status="LETTER_NOT_FOUND" />
            <p>கருப்புப்பெட்டி - எழுத்து சொல்லில் இல்லை</p>
          </div>
        </div>
      </div>
    ),
    target: "#historyboxes",
  },
];

export default steps;
