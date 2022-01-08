// container for storing globals
var store = window.localStorage;
var APP = { wordLength: 5, history: [] };

/**
 * Mottie Keyboard related stuff
 *
 * */
$("#keyboard").keyboard({
  usePreview: false,
  layout: "tamil-tamil99-mod",
  autoAccept: true,
  initialFocus: false,
});

/**
 * Convert a string into an array of strings representing Tamil letters
 *
 * @param {string} word String containing the word to convert to letters
 * @returns an Array of strings that is equivalent to letters in Tamil
 */
function convertToLetters(word) {
  var diacritics = {
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

  var targetList = [];
  for (var i = 0; i != word.length; i++) {
    var ch = word[i];
    diacritics[ch]
      ? (targetList[targetList.length - 1] += ch)
      : targetList.push(ch);
  }
  return targetList;
}

/**
 * Uses the response from the server to populate the UI
 *
 * @param {array} letters Array of string - equivalent to Tamil Letters
 * @param {XMLHttpRequest} http HTTP request object with responseText
 */
function updateResults(letters, http) {
  var status = JSON.parse(http.responseText);
  var box = {
    LETTER_NOT_FOUND: "⬜",
    LETTER_ELSEWHERE: "🟨",
    LETTER_MATCHED: "🟩",
  };

  var statusItem = "";
  var historyItem = $(
    `<div class="is-flex is-justify-content-center my-2"></div>`
  );
  var colors = {
    LETTER_NOT_FOUND: "has-background-light",
    LETTER_ELSEWHERE: "has-background-warning",
    LETTER_MATCHED: "has-background-success",
  };

  for (var i = 0; i < APP.wordLength; i++) {
    var letter = letters[i] || "";
    var color = colors[status[i]] || colors.LETTER_NOT_FOUND;
    statusItem += box[status[i]];
    historyItem.append(
      `<div class="letter-box b-dark ${color}">${letter}</div>`
    );
  }

  APP.history.push(statusItem);
  store.setItem("historyBlocks", JSON.stringify(APP.history));

  $("#historyDiv").append(historyItem);
}

function showError(message) {
  $("#error").text(message).show();
}

function showSuccess() {
  $("#full-history").html(APP.history.join("<br/>"));
  $("#result-modal").addClass("is-active");
}

/**
 * Processes the value in the Input Box
 */
function process() {
  var str = $("#keyboard").val().trim();
  var targetList = convertToLetters(str);

  const http = new XMLHttpRequest();
  http.open("POST", "https://tamilwordle-maleycpqdq-el.a.run.app/verify-word");
  http.setRequestHeader("Content-Type", "application/json");
  http.send(JSON.stringify(targetList));

  http.onreadystatechange = (e) => {
    if (http.readyState === XMLHttpRequest.DONE) {
      switch (http.status) {
        case 202:
          APP.history.push("🟩".repeat(APP.wordLength));
          store.setItem("historyBlocks", JSON.stringify(APP.history));
          showSuccess();
          return;
        case 200:
          updateResults(targetList, http);
          return;
        default:
          showError(http.responseText || "Network Error");
          return;
      }
    }
  };
}

/**
 * Initialize the application
 */
function init() {
  $("#lengthLabel").text(APP.wordLength);
  // See if the game has already been played today
  // If it is already theres, just show the success dialog
  var lastDate = parseInt(store.getItem("date"));
  if (lastDate === new Date().toDateString()) {
    APP.history = JSON.parse(store.getItem("historyBlocks"));
    showSuccess();
    return;
  }

  store.setItem("date", new Date().toDateString());

  const http = new XMLHttpRequest();
  http.open(
    "GET",
    "https://tamilwordle-maleycpqdq-el.a.run.app/get-current-word-len"
  );
  http.onreadystatechange = (e) => {
    if (http.readyState === XMLHttpRequest.DONE) {
      var status = http.status;

      if (status !== 200) {
        console.log("some error happened", http.status);
        showError("Error loading the game!");
        return;
      }

      var jsonResponse = JSON.parse(http.responseText);
      APP.wordLength = jsonResponse.length;
      $("#lengthLabel").text(APP.wordLength);
    }
  };
  http.send();
}

/**
 * Close button clicked on Modal
 */

$(".modal-close").click(function () {
  $(".modal").removeClass("is-active");
});

/**
 * Share Button clicked on Modal
 */
$("#share").click(function () {
  var text = "தமிழ் வோர்டில்\n\n" + APP.history.join("\n");
  if (navigator.share) {
    navigator.share({ text });
  } else {
    navigator.clipboard.writeText(text);
    alert("Content copied to clipboard");
  }
});

window.onload = init();
