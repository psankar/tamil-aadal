package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"unicode"
)

func getWordForToday() string {
	return "கண்ணியம்"
}

const (
	LetterMatched   = "LETTER_MATCHED"
	LetterElseWhere = "LETTER_ELSEWHERE"
	LetterNotFound  = "LETTER_NOT_FOUND"
	UyirMatched     = "UYIR_MATCHED"
	MeiMatched      = "MEI_MATCHED"
)

var todayLetters []string
var todayLettersMap map[string]struct{}
var isDiacritic map[rune]struct{}

var uyirMap, meiMap map[string]string

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

	uyirMap = map[string]string{
		"க": "அ",
		"ங": "அ",
		"ச": "அ",
		"ஞ": "அ",
		"ட": "அ",
		"ண": "அ",
		"த": "அ",
		"ந": "அ",
		"ப": "அ",
		"ம": "அ",
		"ய": "அ",
		"ர": "அ",
		"ல": "அ",
		"வ": "அ",
		"ழ": "அ",
		"ள": "அ",
		"ற": "அ",
		"ன": "அ",

		"கா": "ஆ",
		"ஙா": "ஆ",
		"சா": "ஆ",
		"ஞா": "ஆ",
		"டா": "ஆ",
		"ணா": "ஆ",
		"தா": "ஆ",
		"நா": "ஆ",
		"பா": "ஆ",
		"மா": "ஆ",
		"யா": "ஆ",
		"ரா": "ஆ",
		"லா": "ஆ",
		"வா": "ஆ",
		"ழா": "ஆ",
		"ளா": "ஆ",
		"றா": "ஆ",
		"னா": "ஆ",

		"கி": "இ",
		"ஙி": "இ",
		"சி": "இ",
		"ஞி": "இ",
		"டி": "இ",
		"ணி": "இ",
		"தி": "இ",
		"நி": "இ",
		"பி": "இ",
		"மி": "இ",
		"யி": "இ",
		"ரி": "இ",
		"லி": "இ",
		"வி": "இ",
		"ழி": "இ",
		"ளி": "இ",
		"றி": "இ",
		"னி": "இ",

		"கீ": "ஈ",
		"ஙீ": "ஈ",
		"சீ": "ஈ",
		"ஞீ": "ஈ",
		"டீ": "ஈ",
		"ணீ": "ஈ",
		"தீ": "ஈ",
		"நீ": "ஈ",
		"பீ": "ஈ",
		"மீ": "ஈ",
		"யீ": "ஈ",
		"ரீ": "ஈ",
		"லீ": "ஈ",
		"வீ": "ஈ",
		"ழீ": "ஈ",
		"ளீ": "ஈ",
		"றீ": "ஈ",
		"னீ": "ஈ",

		"கு": "உ",
		"ஙு": "உ",
		"சு": "உ",
		"ஞு": "உ",
		"டு": "உ",
		"ணு": "உ",
		"து": "உ",
		"நு": "உ",
		"பு": "உ",
		"மு": "உ",
		"யு": "உ",
		"ரு": "உ",
		"லு": "உ",
		"வு": "உ",
		"ழு": "உ",
		"ளு": "உ",
		"று": "உ",
		"னு": "உ",

		"கூ": "ஊ",
		"ஙூ": "ஊ",
		"சூ": "ஊ",
		"ஞூ": "ஊ",
		"டூ": "ஊ",
		"ணூ": "ஊ",
		"தூ": "ஊ",
		"நூ": "ஊ",
		"பூ": "ஊ",
		"மூ": "ஊ",
		"யூ": "ஊ",
		"ரூ": "ஊ",
		"லூ": "ஊ",
		"வூ": "ஊ",
		"ழூ": "ஊ",
		"ளூ": "ஊ",
		"றூ": "ஊ",
		"னூ": "ஊ",

		"கெ": "எ",
		"ஙெ": "எ",
		"செ": "எ",
		"ஞெ": "எ",
		"டெ": "எ",
		"ணெ": "எ",
		"தெ": "எ",
		"நெ": "எ",
		"பெ": "எ",
		"மெ": "எ",
		"யெ": "எ",
		"ரெ": "எ",
		"லெ": "எ",
		"வெ": "எ",
		"ழெ": "எ",
		"ளெ": "எ",
		"றெ": "எ",
		"னெ": "எ",

		"கே": "ஏ",
		"ஙே": "ஏ",
		"சே": "ஏ",
		"ஞே": "ஏ",
		"டே": "ஏ",
		"ணே": "ஏ",
		"தே": "ஏ",
		"நே": "ஏ",
		"பே": "ஏ",
		"மே": "ஏ",
		"யே": "ஏ",
		"ரே": "ஏ",
		"லே": "ஏ",
		"வே": "ஏ",
		"ழே": "ஏ",
		"ளே": "ஏ",
		"றே": "ஏ",
		"னே": "ஏ",

		"கை": "ஐ",
		"ஙை": "ஐ",
		"சை": "ஐ",
		"ஞை": "ஐ",
		"டை": "ஐ",
		"ணை": "ஐ",
		"தை": "ஐ",
		"நை": "ஐ",
		"பை": "ஐ",
		"மை": "ஐ",
		"யை": "ஐ",
		"ரை": "ஐ",
		"லை": "ஐ",
		"வை": "ஐ",
		"ழை": "ஐ",
		"ளை": "ஐ",
		"றை": "ஐ",
		"னை": "ஐ",

		"கொ": "ஒ",
		"ஙொ": "ஒ",
		"சொ": "ஒ",
		"ஞொ": "ஒ",
		"டொ": "ஒ",
		"ணொ": "ஒ",
		"தொ": "ஒ",
		"நொ": "ஒ",
		"பொ": "ஒ",
		"மொ": "ஒ",
		"யொ": "ஒ",
		"ரொ": "ஒ",
		"லொ": "ஒ",
		"வொ": "ஒ",
		"ழொ": "ஒ",
		"ளொ": "ஒ",
		"றொ": "ஒ",
		"னொ": "ஒ",

		"கோ": "ஓ",
		"ஙோ": "ஓ",
		"சோ": "ஓ",
		"ஞோ": "ஓ",
		"டோ": "ஓ",
		"ணோ": "ஓ",
		"தோ": "ஓ",
		"நோ": "ஓ",
		"போ": "ஓ",
		"மோ": "ஓ",
		"யோ": "ஓ",
		"ரோ": "ஓ",
		"லோ": "ஓ",
		"வோ": "ஓ",
		"ழோ": "ஓ",
		"ளோ": "ஓ",
		"றோ": "ஓ",
		"னோ": "ஓ",

		"கௌ": "ஔ",
		"ஙௌ": "ஔ",
		"சௌ": "ஔ",
		"ஞௌ": "ஔ",
		"டௌ": "ஔ",
		"ணௌ": "ஔ",
		"தௌ": "ஔ",
		"நௌ": "ஔ",
		"பௌ": "ஔ",
		"மௌ": "ஔ",
		"யௌ": "ஔ",
		"ரௌ": "ஔ",
		"லௌ": "ஔ",
		"வௌ": "ஔ",
		"ழௌ": "ஔ",
		"ளௌ": "ஔ",
		"றௌ": "ஔ",
		"னௌ": "ஔ",

		"அ": "அ",
		"ஆ": "ஆ",
		"இ": "இ",
		"ஈ": "ஈ",
		"உ": "உ",
		"ஊ": "ஊ",
		"எ": "எ",
		"ஏ": "ஏ",
		"ஐ": "ஐ",
		"ஒ": "ஒ",
		"ஓ": "ஓ",
		"ஔ": "ஔ",
	}

	meiMap = map[string]string{
		"க்": "க்",
		"ங்": "ங்",
		"ச்": "ச்",
		"ஞ்": "ஞ்",
		"ட்": "ட்",
		"ண்": "ண்",
		"த்": "த்",
		"ந்": "ந்",
		"ப்": "ப்",
		"ம்": "ம்",
		"ய்": "ய்",
		"ர்": "ர்",
		"ல்": "ல்",
		"வ்": "வ்",
		"ழ்": "ழ்",
		"ள்": "ள்",
		"ற்": "ற்",
		"ன்": "ன்",

		"க":  "க்",
		"கா": "க்",
		"கி": "க்",
		"கீ": "க்",
		"கு": "க்",
		"கூ": "க்",
		"கெ": "க்",
		"கே": "க்",
		"கை": "க்",
		"கொ": "க்",
		"கோ": "க்",
		"கௌ": "க்",

		"ங":  "ங்",
		"ஙா": "ங்",
		"ஙி": "ங்",
		"ஙீ": "ங்",
		"ஙு": "ங்",
		"ஙூ": "ங்",
		"ஙெ": "ங்",
		"ஙே": "ங்",
		"ஙை": "ங்",
		"ஙொ": "ங்",
		"ஙோ": "ங்",
		"ஙௌ": "ங்",

		"ச":  "ச்",
		"சா": "ச்",
		"சி": "ச்",
		"சீ": "ச்",
		"சு": "ச்",
		"சூ": "ச்",
		"செ": "ச்",
		"சே": "ச்",
		"சை": "ச்",
		"சொ": "ச்",
		"சோ": "ச்",
		"சௌ": "ச்",

		"ஞ":  "ஞ்",
		"ஞா": "ஞ்",
		"ஞி": "ஞ்",
		"ஞீ": "ஞ்",
		"ஞு": "ஞ்",
		"ஞூ": "ஞ்",
		"ஞெ": "ஞ்",
		"ஞே": "ஞ்",
		"ஞை": "ஞ்",
		"ஞொ": "ஞ்",
		"ஞோ": "ஞ்",
		"ஞௌ": "ஞ்",

		"ட":  "ட்",
		"டா": "ட்",
		"டி": "ட்",
		"டீ": "ட்",
		"டு": "ட்",
		"டூ": "ட்",
		"டெ": "ட்",
		"டே": "ட்",
		"டை": "ட்",
		"டொ": "ட்",
		"டோ": "ட்",
		"டௌ": "ட்",

		"ண":  "ண்",
		"ணா": "ண்",
		"ணி": "ண்",
		"ணீ": "ண்",
		"ணு": "ண்",
		"ணூ": "ண்",
		"ணெ": "ண்",
		"ணே": "ண்",
		"ணை": "ண்",
		"ணொ": "ண்",
		"ணோ": "ண்",
		"ணௌ": "ண்",

		"த":  "த்",
		"தா": "த்",
		"தி": "த்",
		"தீ": "த்",
		"து": "த்",
		"தூ": "த்",
		"தெ": "த்",
		"தே": "த்",
		"தை": "த்",
		"தொ": "த்",
		"தோ": "த்",
		"தௌ": "த்",

		"ந":  "ந்",
		"நா": "ந்",
		"நி": "ந்",
		"நீ": "ந்",
		"நு": "ந்",
		"நூ": "ந்",
		"நெ": "ந்",
		"நே": "ந்",
		"நை": "ந்",
		"நொ": "ந்",
		"நோ": "ந்",
		"நௌ": "ந்",

		"ப":  "ப்",
		"பா": "ப்",
		"பி": "ப்",
		"பீ": "ப்",
		"பு": "ப்",
		"பூ": "ப்",
		"பெ": "ப்",
		"பே": "ப்",
		"பை": "ப்",
		"பொ": "ப்",
		"போ": "ப்",
		"பௌ": "ப்",

		"ம":  "ம்",
		"மா": "ம்",
		"மி": "ம்",
		"மீ": "ம்",
		"மு": "ம்",
		"மூ": "ம்",
		"மெ": "ம்",
		"மே": "ம்",
		"மை": "ம்",
		"மொ": "ம்",
		"மோ": "ம்",
		"மௌ": "ம்",

		"ய":  "ய்",
		"யா": "ய்",
		"யி": "ய்",
		"யீ": "ய்",
		"யு": "ய்",
		"யூ": "ய்",
		"யெ": "ய்",
		"யே": "ய்",
		"யை": "ய்",
		"யொ": "ய்",
		"யோ": "ய்",
		"யௌ": "ய்",

		"ர":  "ர்",
		"ரா": "ர்",
		"ரி": "ர்",
		"ரீ": "ர்",
		"ரு": "ர்",
		"ரூ": "ர்",
		"ரெ": "ர்",
		"ரே": "ர்",
		"ரை": "ர்",
		"ரொ": "ர்",
		"ரோ": "ர்",
		"ரௌ": "ர்",

		"ல":  "ல்",
		"லா": "ல்",
		"லி": "ல்",
		"லீ": "ல்",
		"லு": "ல்",
		"லூ": "ல்",
		"லெ": "ல்",
		"லே": "ல்",
		"லை": "ல்",
		"லொ": "ல்",
		"லோ": "ல்",
		"லௌ": "ல்",

		"வ":  "வ்",
		"வா": "வ்",
		"வி": "வ்",
		"வீ": "வ்",
		"வு": "வ்",
		"வூ": "வ்",
		"வெ": "வ்",
		"வே": "வ்",
		"வை": "வ்",
		"வொ": "வ்",
		"வோ": "வ்",
		"வௌ": "வ்",

		"ழ":  "ழ்",
		"ழா": "ழ்",
		"ழி": "ழ்",
		"ழீ": "ழ்",
		"ழு": "ழ்",
		"ழூ": "ழ்",
		"ழெ": "ழ்",
		"ழே": "ழ்",
		"ழை": "ழ்",
		"ழொ": "ழ்",
		"ழோ": "ழ்",
		"ழௌ": "ழ்",

		"ள":  "ள்",
		"ளா": "ள்",
		"ளி": "ள்",
		"ளீ": "ள்",
		"ளு": "ள்",
		"ளூ": "ள்",
		"ளெ": "ள்",
		"ளே": "ள்",
		"ளை": "ள்",
		"ளொ": "ள்",
		"ளோ": "ள்",
		"ளௌ": "ள்",

		"ற":  "ற்",
		"றா": "ற்",
		"றி": "ற்",
		"றீ": "ற்",
		"று": "ற்",
		"றூ": "ற்",
		"றெ": "ற்",
		"றே": "ற்",
		"றை": "ற்",
		"றொ": "ற்",
		"றோ": "ற்",
		"றௌ": "ற்",

		"ன":  "ன்",
		"னா": "ன்",
		"னி": "ன்",
		"னீ": "ன்",
		"னு": "ன்",
		"னூ": "ன்",
		"னெ": "ன்",
		"னே": "ன்",
		"னை": "ன்",
		"னொ": "ன்",
		"னோ": "ன்",
		"னௌ": "ன்",
	}
}

type CurrentWordLenResponse struct {
	Length int
}

func getCurrentWordLenHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	err := json.NewEncoder(w).Encode(CurrentWordLenResponse{len(todayLetters)})
	if err != nil {
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func verifyWordHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	var letters []string
	err := json.NewDecoder(r.Body).Decode(&letters)
	if err != nil {
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	if len(letters) != len(todayLetters) {
		http.Error(w, "Invalid word length; சரியான நீளத்தில் அனுப்பவும்",
			http.StatusBadRequest)
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
		w.Write([]byte("OK"))
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func verifyWordWithUyirMeiHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	var letters []string
	err := json.NewDecoder(r.Body).Decode(&letters)
	if err != nil {
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	if len(letters) != len(todayLetters) {
		http.Error(w, "Invalid word length; சரியான நீளத்தில் அனுப்பவும்",
			http.StatusBadRequest)
		return
	}

	response, allMatched := verifyWordWithUyirMei(letters, todayLetters)

	w.Header().Set("Content-Type", "application/json")
	if allMatched {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("OK"))
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func verifyWordWithUyirMei(gotLetters []string,
	wantLetters []string) ([][]string, bool) {

	allMatched := true

	var response [][]string
	for i, letter := range gotLetters {
		var curLetterResponse []string
		if letter == wantLetters[i] {
			curLetterResponse = []string{LetterMatched}
			response = append(response, curLetterResponse)
			continue
		}

		allMatched = false

		if _, found := todayLettersMap[letter]; found {
			curLetterResponse = []string{LetterElseWhere}
		} else {
			curLetterResponse = []string{LetterNotFound}
		}

		targetUyir, ok := uyirMap[wantLetters[i]]
		if ok && (targetUyir == uyirMap[letter]) {
			curLetterResponse = append(curLetterResponse, UyirMatched)
		}

		targetMei, ok := meiMap[wantLetters[i]]
		if ok && (targetMei == meiMap[letter]) {
			curLetterResponse = append(curLetterResponse, MeiMatched)
		}

		response = append(response, curLetterResponse)
	}

	return response, allMatched
}

func splitWordGetLetters(word string) ([]string, error) {
	var letters []string

	for _, r := range word {
		if !unicode.Is(unicode.Tamil, r) {
			return nil, fmt.Errorf("Non-Tamil word")
		}

		if _, yes := isDiacritic[r]; yes {
			if len(letters) == 0 {
				return nil, fmt.Errorf("invalid diacritic position")
			}
			letters[len(letters)-1] += string(r)
		} else {
			letters = append(letters, string(r))
		}
	}

	return letters, nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui1", http.StatusSeeOther)
}

func enableCORS(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding")
}

func main() {
	log.Print("starting server...")

	http.HandleFunc("/get-current-word-len", getCurrentWordLenHandler)
	http.HandleFunc("/verify-word", verifyWordHandler)
	http.HandleFunc("/verify-word-with-uyirmei", verifyWordWithUyirMeiHandler)

	ui1 := http.FileServer(http.Dir("./ui1"))
	ui2 := http.FileServer(http.Dir("./ui2"))
	ui3 := http.FileServer(http.Dir("./ui3"))
	http.Handle("/ui1/", http.StripPrefix("/ui1/", ui1))
	http.Handle("/ui2/", http.StripPrefix("/ui2/", ui2))
	http.Handle("/ui3/", http.StripPrefix("/ui3/", ui3))

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
