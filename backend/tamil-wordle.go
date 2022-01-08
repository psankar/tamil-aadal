package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"unicode"
)

const (
	LetterMatched   = "LETTER_MATCHED"
	LetterElseWhere = "LETTER_ELSEWHERE"
	LetterNotFound  = "LETTER_NOT_FOUND"
)

// Allowed Origins for CORS requests, initialized in init
var allowedOrigins map[string]bool

var todayLetters []string
var todayLettersMap map[string]struct{}
var isDiacritic map[rune]struct{}
var tmpl *template.Template

func init() {
	var empty struct{}
	isDiacritic = make(map[rune]struct{})
	isDiacritic['\u0B82'] = empty
	isDiacritic['\u0BBE'] = empty
	isDiacritic['\u0BBF'] = empty
	isDiacritic['\u0BC0'] = empty
	isDiacritic['\u0BC1'] = empty
	isDiacritic['\u0BC2'] = empty
	isDiacritic['\u0BC6'] = empty
	isDiacritic['\u0BC7'] = empty
	isDiacritic['\u0BC8'] = empty
	isDiacritic['\u0BCA'] = empty
	isDiacritic['\u0BCB'] = empty
	isDiacritic['\u0BCC'] = empty
	isDiacritic['\u0BCD'] = empty
	isDiacritic['\u0BD7'] = empty

	var err error
	todayLetters, err = splitWordGetLetters(getWordForToday())
	if err != nil {
		log.Fatal(err)
		return
	}

	todayLettersMap = make(map[string]struct{})
	for _, letter := range todayLetters {
		todayLettersMap[letter] = empty
	}

	tmpl, err = template.New("wordle").Parse(htmlFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	// initialize allowed origins
	allowedOrigins = map[string]bool{
		"https://tamil-wordle.local.vercel.app":        true,
		"https://tamil-wordle.vercel.app":              true,
		"https://tamil-wordle-tsureshkumar.vercel.app": true,
		"https://tecoholic.github.io":                  true,
	}
}

type CurrentWordLenResponse struct {
	Length int
}

func getWordForToday() string {
	return "தமிழ்"
}

func getCurrentWordLenHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(CurrentWordLenResponse{len(todayLetters)})
	if err != nil {
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func verifyWordHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r, r.Header.Get("Origin"))
	if r.Method == "OPTIONS" {
		return
	}
	var letters []string
	err := json.NewDecoder(r.Body).Decode(&letters)
	if err != nil {
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்", http.StatusBadRequest)
		return
	}

	if len(letters) != len(todayLetters) {
		http.Error(w, "Invalid word length; சரியான நீளத்தில் அனுப்பவும்", http.StatusBadRequest)
		return
	}

	allMatched := true

	var response []string
	for i := 0; i < len(letters); i++ {
		// log.Printf("DEBUG: %q == %q", letters[i], todayLetters[i])
		if letters[i] == todayLetters[i] {
			response = append(response, LetterMatched)
		} else {
			allMatched = false
			if _, found := todayLettersMap[letters[i]]; found {
				response = append(response, LetterElseWhere)
			} else {
				response = append(response, LetterNotFound)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")

	if allMatched {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func splitWordGetLetters(word string) ([]string, error) {
	var letters []string

	for _, r := range word {
		if !unicode.Is(unicode.Tamil, r) {
			return nil, fmt.Errorf("Non-Tamil word")
		}

		if _, yes := isDiacritic[r]; yes {
			if len(letters) == 0 {
				return nil, fmt.Errorf("Invalid diacritic position")
			}
			letters[len(letters)-1] += string(r)
		} else {
			letters = append(letters, string(r))
		}
	}

	return letters, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func enableCORS(w http.ResponseWriter, req *http.Request, origin string) {
	if allowedOrigins[origin] {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

func main() {
	log.Print("starting server...")

	http.HandleFunc("/get-current-word-len", getCurrentWordLenHandler)
	http.HandleFunc("/verify-word", verifyWordHandler)
	http.HandleFunc("/", homeHandler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

const htmlFile = `<html>
<meta charset="UTF-8" />

<head>
  <title>Tamil Wordle</title>
</head>

<body>
  <label id="lengthLabel"></label> எழுத்து(க்)கள் அளவு நீளமான தமிழ்ச்சொல்லை,
  தமிழில் தட்டச்சு செய்து, 'சரி' பொத்தானை அழுத்தவும் <br />
  <input type="text" id="curword" />
  <button onclick="process()">சரி</button>
  <hr />
  <p>❤️ - சரியான எழுத்து</p>
  <p>&#128584; - இல்லாத எழுத்து, கடல்லையே இல்லையாம்</p>
  <p>&#128064; - தவறான இடத்தில் உள்ள சரியான எழுத்து</p>
  <hr />
  <div id="tilesDiv"></div>
  <hr />
  <div id="historyDiv"></div>
</body>

<script>
  function process() {
	var str = document.getElementById("curword").value.trim();

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
	for (var i = 0; i != str.length; i++) {
	  var ch = str[i];
	  diacritics[ch]
		? (targetList[targetList.length - 1] += ch)
		: targetList.push(ch);
	}

	const http = new XMLHttpRequest();
	http.open("POST", "/verify-word");
	http.setRequestHeader("Content-Type", "application/json");
	http.send(JSON.stringify(targetList));

	http.onreadystatechange = (e) => {
	  if (http.readyState === XMLHttpRequest.DONE) {
		switch (http.status) {
		  case 202:
			alert(
			  "சரியான சொல்லைக் கண்டுபிடித்துவிட்டீர்கள் !!! If you are interested, copy and paste the emoji table to social media."
			);
			var tilesDiv = document.getElementById("tilesDiv");
			var newLabel = document.createElement("Label");
			for (var i = 0; i < jsonResponse.length; i++) {
			  newLabel.innerHTML += " ❤️ ";
			}
			tilesDiv.appendChild(newLabel);

			return;
		  case 200:
			jsonResponse = JSON.parse(http.responseText);

			var historyDiv = document.getElementById("historyDiv");
			var historyEntry = document.createElement("Label");
			historyEntry.innerHTML = str;
			var historyBreak = document.createElement("br");
			historyDiv.appendChild(historyEntry);
			historyDiv.appendChild(historyBreak);

			var tilesDiv = document.getElementById("tilesDiv");
			var newLabel = document.createElement("Label");

			for (var i = 0; i < jsonResponse.length; i++) {
			  var resp = jsonResponse[i];
			  switch (resp) {
				case "LETTER_NOT_FOUND":
				  newLabel.innerHTML += " &#128584; ";
				  break;
				case "LETTER_ELSEWHERE":
				  newLabel.innerHTML += " &#128064; ";
				  break;
				case "LETTER_MATCHED":
				  newLabel.innerHTML += " ❤️ ";
				  break;
				default:
				  alert("Error in game:", resp);
				  break;
			  }
			}
			tilesDiv.appendChild(newLabel);
			tilesBreak = document.createElement("br");
			tilesDiv.appendChild(tilesBreak);
			return;
		  default:
			alert(http.responseText);
			return;
		}
	  }
	};
  }

  function loadLen() {
	const http = new XMLHttpRequest();
	http.open("GET", "/get-current-word-len");
	http.onreadystatechange = (e) => {
	  if (http.readyState === XMLHttpRequest.DONE) {
		var status = http.status;

		if (status !== 200) {
		  console.log("some error happened", http.status);
		  alert("Error loading the game!");
		  return;
		}

		var jsonResponse = JSON.parse(http.responseText);
		var lengthLabel = document.getElementById("lengthLabel");
		lengthLabel.innerText = jsonResponse.Length;
	  }
	};
	http.send();
  }

  window.onload = loadLen();
</script>
</html>
`
