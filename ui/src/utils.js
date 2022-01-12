export const diacritics = {
  "\u0B82": true,
  "\u0BBE": true,
  "\u0BBF": true,
  "\u0BC0": true,
  "\u0BC1": true,
  "\u0BC2": true,
  "\u0BC6": true,
  "\u0BC7": true,
  "\u0BC8": true,
  "\u0BCA": true,
  "\u0BCB": true,
  "\u0BCC": true,
  "\u0BCD": true,
  "\u0BD7": true,
};

export const vowels = {
  "\u0B85": true,
  "\u0B86": true,
  "\u0B87": true,
  "\u0B88": true,
  "\u0B89": true,
  "\u0B8A": true,
  "\u0B8E": true,
  "\u0B8F": true,
  "\u0B90": true,
  "\u0B92": true,
  "\u0B93": true,
  "\u0B94": true,
};

/**
 * Convert a string into an array of strings representing Tamil letters
 *
 * @param {string} word String containing the word to convert to letters
 * @returns an Array of strings that is equivalent to letters in Tamil
 */
export function toTamilLetters(word) {
  let letters = [];
  for (let i = 0; i !== word.length; i++) {
    let ch = word[i];
    diacritics[ch] && letters.length && letters[letters.length - 1].length < 2
      ? (letters[letters.length - 1] += ch)
      : letters.push(ch);
  }
  return letters;
}

export const PAGES = {
  INSTRUCTIONS: "instructions",
  SETTINGS: "settings",
  WORKBENCH: "workbench",
  SUCCESS: "success",
};
