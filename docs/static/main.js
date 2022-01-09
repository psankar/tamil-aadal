// container for storing globals
var store = window.localStorage;
var APP = { wordLength: 5, history: [] };
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

/**
 * Mottie Keyboard related stuff
 *
 * */
function initKeyboard() {
  $("#keyboard").val("");
  $("#keyboard").keyboard({
    usePreview: false,
    layout: "tamil-tamil99-mod",
    autoAccept: true,
    alwaysOpen: true,
    appendTo: "#keyboard-container",
    change: function () {
      renderWorkarea();
    },
  });
}

/**
 * Convert a string into an array of strings representing Tamil letters
 *
 * @param {string} word String containing the word to convert to letters
 * @returns an Array of strings that is equivalent to letters in Tamil
 */
function convertToLetters(word) {
  var targetList = [];
  for (var i = 0; i != word.length; i++) {
    var ch = word[i];
    diacritics[ch] && targetList[targetList.length - 1].length < 2
      ? (targetList[targetList.length - 1] += ch)
      : targetList.push(ch);
  }
  return targetList;
}

/**
 * Renders the boxes which will show the typing results
 */
function renderWorkarea() {
  var word = $("#keyboard").val();
  var letters = convertToLetters(word).slice(0, APP.wordLength);
  $("#keyboard").val(letters.join(""));
  var workarea = $("#workarea");
  workarea.empty();

  for (var i = 0; i < APP.wordLength; i++) {
    var _class = "letter-box" + (letters[i] ? " b-dark" : "");
    workarea.append(`<div class="${_class}">${letters[i] || ""}</div>`);
  }
  if (letters.length === APP.wordLength) {
    $("#verify-button").show();
  } else {
    $("#verify-button").hide();
  }
}

/**
 * Uses the response from the server to populate the UI
 *
 * @param {array} letters Array of string - equivalent to Tamil Letters
 * @param {XMLHttpRequest} http HTTP request object with responseText
 */
function updateResults(letters, status, success) {
  var box = {
    LETTER_NOT_FOUND: "⚫",
    LETTER_ELSEWHERE: "🟡",
    LETTER_MATCHED: "🟢",
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
  $("#keyboard").val("");
  if (success) {
    showSuccess();
  } else {
    renderWorkarea();
  }
}

function showError(message) {
  $("#error").text(message).show();
}

function showSuccess() {
  $("#full-history").html(APP.history.join("<br/>"));
  $("#result-modal").addClass("is-active");
  store.setItem("successdate", new Date().toDateString());
}

/**
 * Processes the value in the Input Box
 */
function process() {
  $("#error").hide();
  var word = $("#keyboard").val();
  var letters = convertToLetters(word).slice(0, APP.wordLength);
  $.ajax({
    method: "POST",
    url: "https://tamilwordle-maleycpqdq-el.a.run.app/verify-word",
    data: JSON.stringify(letters),
    dataType: "text",
    contentType: "application/json",
    success: function (data, status, xhr) {
      switch (xhr.status) {
        case 202:
          updateResults(
            letters,
            Array(APP.wordLength).fill("LETTER_MATCHED"),
            true
          );
          return;
        case 200:
          updateResults(letters, JSON.parse(data));
          return;
        default:
          showError(http.responseText || "Network Error");
          return;
      }
    },
  }).fail(function () {
    showError("வார்த்தையைச் சரி பார்க்க முடியவில்லை");
    renderWorkarea();
  });
  $("#keybaord").val("");
  $("#verify-button").hide();
}

/**
 * Initialize the application
 */
function init() {
  // See if the game has already been played today
  // If it is already there, just show the success dialog
  if (store.getItem("successdate") === new Date().toDateString()) {
    APP.history = JSON.parse(store.getItem("historyBlocks"));
    showSuccess();
    return;
  }
  initKeyboard();
  $("#verify-button").hide();

  $.ajax({
    url: "https://tamilwordle-maleycpqdq-el.a.run.app/get-current-word-len",
    success: function (data) {
      APP.wordLength = JSON.parse(data).Length;
      $("#lengthLabel").text(APP.wordLength);
      var workarea = $("#workarea");
      for (var i = 0; i < APP.wordLength; i++) {
        workarea.append(`<div class="letter-box"></div>`);
      }
      $("#verify-button").css({
        width: (64 + 4) * APP.wordLength + 0.75 * (APP.wordLength - 1) + "px",
      });
    },
  }).fail(function () {
    console.log("some error happened", http.status);
    showError("Error loading the game!");
  });
}

// Check for click events on the navbar burger icon
$(".navbar-burger").click(function () {
  // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
  $(".navbar-burger").toggleClass("is-active");
  $(".navbar-menu").toggleClass("is-active");
});

//Close button clicked on Modal
$(".modal-close").click(function () {
  $(".modal").removeClass("is-active");
});

// Share Button clicked on Modal
$("#share").click(function () {
  var d = new Date();
  var text =
    `தமிழ் வோர்டில் (${d.getDate()}/${d.getMonth() + 1})\n\n` +
    APP.history.join("\n");
  if (navigator.share) {
    navigator.share({ text });
  } else {
    navigator.clipboard.writeText(text);
    alert("Content copied to clipboard");
  }
});

window.onload = init();
