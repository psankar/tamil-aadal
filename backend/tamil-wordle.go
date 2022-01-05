package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"unicode"
)

const (
	LetterMatched   = "LETTER_MATCHED"
	LetterElseWhere = "LETTER_ELSEWHERE"
	LetterNotFound  = "LETTER_NOT_FOUND"
)

var todayLetters []string
var todayLettersMap map[string]struct{}
var isDiacritic map[rune]struct{}

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
		log.Printf("DEBUG: %q == %q", letters[i], todayLetters[i])
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

func main() {
	http.HandleFunc("/get-current-word-len", getCurrentWordLenHandler)
	http.HandleFunc("/verify-word", verifyWordHandler)
	http.ListenAndServe(":8080", nil)
}
