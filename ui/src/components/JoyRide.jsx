import Cell from "./Cell";

const steps = [
  {
    content: (
      <div>
        <h2 className="has-text-weight-bold mb-3">
          இந்த விளையாட்டு விளையாடுவது எப்படி?
        </h2>
        <p>தினம் ஒரு புது வார்த்தைய நீங்கள் கண்டுபிடிக்க வேண்டும்</p>
      </div>
    ),
    placement: "center",
    target: "body",
  },
  {
    content: (
      <h3>
        அந்த வார்த்தையில் எத்தனை எழுத்துக்கள் உள்ளனவே, அவ்வளவு பெட்டிகள்
        விளையாட்டுத் தளத்தில் இருக்கும். அத்தனை எழுத்துக்கள் உள்ள எதேனும் ஒரு
        வார்த்தையை யோசித்துக் கொள்ளுங்கள்.
      </h3>
    ),
    spotlightPadding: 4,
    target: "#input-boxes",
  },
  {
    content: (
      <p>
        நீங்கள் யூகித்த வார்த்தையைக் கீழே இருக்கும் தட்டச்சுப்பலகை மூலம்
        உள்ளிடவும்.
      </p>
    ),
    target: "#keyboard",
  },
  {
    content: <p>வார்த்தையை உள்ளிட்டவுடன், "சரிபார்" பொத்தானை அழுத்தவும்</p>,
    target: "#verify-button",
  },
  {
    content: (
      <div>
        <p>
          இன்றய வார்த்தையின் எழுத்துக்களும் நீங்கள் உள்ளிட்ட வார்த்தையில் உள்ள
          எழுத்துக்களும், எப்படிப் பொருந்தியுள்ளது என்று பொட்டிகளின் வண்ணங்களைக்
          கொண்டுத் தெரிந்து கொள்ளலாம்.
        </p>
        <div className="my-3">
          <div className="is-flex is-justify-content-center">
            <Cell letter="" status="LETTER_MATCHED" />
            <p>பெட்டி பச்சையாக மாறினால், எழுத்து பொருத்தமான இடத்தில் உள்ளது.</p>
          </div>
        </div>
        <div className="my-3">
          <div className="is-flex is-justify-content-center">
            <Cell letter="" status="LETTER_ELSEWHERE" />
            <p>பெட்டி மஞ்சளாக மாறினால், எழுத்து தவறான் இடத்தில் உள்ளது.</p>
          </div>
        </div>

        <div className="my-3">
          <div className="is-flex is-justify-content-center">
            <Cell letter="" status="LETTER_NOT_FOUND" />
            <p>
              பெட்டி இளங்கருப்பாக மாறினால், எழுத்து இன்றய வார்த்தையில் இடம்
              பிடிக்கவில்லை.
            </p>
          </div>
        </div>
      </div>
    ),
    target: "#historyboxes",
  },
];

export default steps;
