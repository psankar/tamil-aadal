package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"tamilaadal.com/backend/dao"
)

const (
	LetterMatched   = "LETTER_MATCHED"
	LetterElseWhere = "LETTER_ELSEWHERE"
	LetterNotFound  = "LETTER_NOT_FOUND"
	UyirMatched     = "UYIR_MATCHED"
	MeiMatched      = "MEI_MATCHED"
)

var uyirMap, meiMap map[string]string

const pubKeyPath = "auth/admin.rsa.pub"
const privKeyPath = "auth/admin.rsa"

const (
	Issuer   = "tamilaadal-admin"
	Audience = "tamilaadal"
)

var verifyKey *rsa.PublicKey
var signKey *rsa.PrivateKey

func init() {
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

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
		return
	}

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key: %s", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("failed to parse private key: %s", err)
	}
}

type WordMetaResponse struct {
	Length int
	User   dao.User
}

func getWordMetaHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method == http.MethodOptions {
		return
	}

	// check if date param is present
	date := r.URL.Query().Get("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	wordWrapper, err := dao.GetWordForTheDay(date)
	if err != nil || wordWrapper.Word.Id == "" {
		log.Printf("failed to get word for the day: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}

	wantLetters := wordWrapper.Letters

	err = json.NewEncoder(w).Encode(WordMetaResponse{len(wantLetters), wordWrapper.Word.User})
	if err != nil {
		log.Printf("failed to encode response: %s", err)
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

	// check if date param is present
	date := r.URL.Query().Get("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	wordWrapper, err := dao.GetWordForTheDay(date)
	if err != nil || wordWrapper.Word.Id == "" {
		log.Printf("failed to get word for the day: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}

	wantLetters := wordWrapper.Letters
	wantLettersMap := wordWrapper.LettersMap

	if len(letters) != len(wantLetters) {
		http.Error(w, "Invalid word length; சரியான நீளத்தில் அனுப்பவும்",
			http.StatusBadRequest)
		return
	}

	allMatched := true

	var response []string
	for i := 0; i < len(letters); i++ {
		// log.Printf("DEBUG: %q == %q", letters[i], todayLetters[i])
		if letters[i] == wantLetters[i] {
			response = append(response, LetterMatched)
		} else {
			allMatched = false
			if _, found := wantLettersMap[letters[i]]; found {
				response = append(response, LetterElseWhere)
			} else {
				response = append(response, LetterNotFound)
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")

	if allMatched {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("ok"))
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode response: %s", err)
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

	// check if date param is present
	date := r.URL.Query().Get("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	wordWrapper, err := dao.GetWordForTheDay(date)
	if err != nil || wordWrapper.Word.Id == "" {
		log.Printf("failed to get word for the day: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}

	wantLetters := wordWrapper.Letters
	wantLettersMap := wordWrapper.LettersMap

	if len(letters) != len(wantLetters) {
		http.Error(w, "Invalid word length; சரியான நீளத்தில் அனுப்பவும்",
			http.StatusBadRequest)
		return
	}

	response, allMatched := verifyWordWithUyirMei(letters, wantLetters, wantLettersMap)

	w.Header().Set("Content-Type", "application/json")
	if allMatched {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("OK"))
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("failed to encode response: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func verifyWordWithUyirMei(gotLetters []string,
	wantLetters []string, wantLettersMap map[string]struct{}) ([][]string, bool) {
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

		if _, found := wantLettersMap[letter]; found {
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/ui1", http.StatusSeeOther)
}

func generateAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	var userID string
	err := json.NewDecoder(r.Body).Decode(&userID)
	if err != nil {
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	// create a signer for rsa 512
	t := jwt.New(jwt.GetSigningMethod("RS512"))

	// set our claims
	t.Claims =
		&jwt.StandardClaims{
			// set the expire time
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    Issuer,
			Audience:  Audience,
			Subject:   userID,
		}

	// Create token string
	token, err := t.SignedString(signKey)
	if err != nil {
		log.Printf("failed to generate token: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func keyGenHandler(w http.ResponseWriter, r *http.Request) {
	// extract token from request params
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Invalid token; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	// parse token
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	var claims jwt.MapClaims
	var ok bool
	if claims, ok = t.Claims.(jwt.MapClaims); !ok || !t.Valid {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
	// valid sub claim from JWT with the request body
	userId := claims["sub"].(string)

	// Get user from request params
	u := r.URL.Query().Get("user")
	if userId != u {
		http.Error(w, "Unauthorized; உங்கள் புகுபதிகை தவறானது",
			http.StatusUnauthorized)
		return
	}

	// Generate a key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Printf("failed to generate key pair: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
	publicKey := privateKey.Public()
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey.(*rsa.PublicKey)),
	})

	// Update public key in DB
	err = dao.UpdatePublicKey(userId, string(pubKeyPEM))
	if err != nil {
		log.Printf("failed to update public key: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}

	// Return private key
	w.WriteHeader(http.StatusOK)

	// Return a HTML page with the private key
	// read HTML template
	tpl, err := template.ParseFiles("ui1/magic.html")
	if err != nil {
		log.Printf("failed to parse template: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
	// execute template
	user, err := dao.GetUser(userId)
	if err != nil {
		log.Printf("failed to get user: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
	var values = map[string]string{
		"PrivateKey":    string(keyPEM),
		"UserId":        userId,
		"UserName":      user.Name,
		"TwitterHandle": user.TwitterHandle,
	}

	err = tpl.Execute(w, values)
	if err != nil {
		log.Printf("failed to execute template: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
}

func addWordHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from request body
	var u map[string]string
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	var word dao.Word
	word.Word = u["word"]
	word.Date = u["date"]

	var user dao.User
	user.Id = u["userId"]
	user.Name = u["userName"]
	user.TwitterHandle = u["twitterHandle"]
	word.User = user

	id, err := dao.AddWord(word)
	if err != nil {
		log.Printf("failed to add word: %s", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்\n"+err.Error(),
			http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Word added successfully with id: " + id))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from request body
	var u dao.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println("failed to decode user: ", err)
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	// Create user
	id, err := dao.CreateUser(u)
	if err != nil {
		log.Println("failed to create user: ", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}

func markUserActiveHandler(w http.ResponseWriter, r *http.Request) {
	// Get user from request body
	var userId string
	err := json.NewDecoder(r.Body).Decode(&userId)
	if err != nil {
		log.Println("failed to decode userId: ", err)
		http.Error(w, "Invalid body; தப்புதப்பா அனுப்ப வேண்டாம்",
			http.StatusBadRequest)
		return
	}

	// Mark user active
	err = dao.MarkUserActive(userId)
	if err != nil {
		log.Println("failed to mark user active: ", err)
		http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
			http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User marked active successfully"))
}

// jwt.ParseRSAPublicKeyFromPEM has a bug in it, so we need to do it ourselves
// https://github.com/golang-jwt/jwt/issues/119
// ParseRSAPublicKeyFromPEM parses a PEM encoded PKCS1 public key
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, fmt.Errorf("key must be PEM encoded")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}
	return parsedKey.(*rsa.PublicKey), nil
}

func jwtAdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				next.ServeHTTP(w, r)
			} else {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func jwtUserAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		// Get user from header
		userId := r.Header.Get("x-user-id")

		var key *rsa.PublicKey

		// Use user key for validation
		if userId != "" {
			user, err := dao.GetUser(userId)
			if err != nil {
				log.Println("Error getting user: ", err)
				http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
					http.StatusInternalServerError)
				return
			}
			if !user.Active {
				log.Println("User is not active but is trying to set the word")
				http.Error(w, "User is not active", http.StatusUnauthorized)
				return
			}
			key, err = ParseRSAPublicKeyFromPEM([]byte(user.PublicKey))

			if err != nil {
				log.Println("Error parsing user public key: ", err)
				http.Error(w, "Internal error; தடங்கலுக்கு வருந்துகிறோம்",
					http.StatusInternalServerError)
				return
			}
		} else {
			// user id is not present. unauthorized
			log.Println("User id not present in header")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				return key, nil
			})

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				next.ServeHTTP(w, r)
			} else {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}

func enableCORS(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding")
}

func main() {
	log.Print("starting server...")

	http.HandleFunc("/get-word-meta", getWordMetaHandler)
	http.HandleFunc("/verify-word", verifyWordHandler)
	http.HandleFunc("/verify-word-with-uyirmei", verifyWordWithUyirMeiHandler)

	ui1 := http.FileServer(http.Dir("./ui1"))
	ui2 := http.FileServer(http.Dir("./ui2"))
	ui3 := http.FileServer(http.Dir("./ui3"))
	http.Handle("/ui1/", http.StripPrefix("/ui1/", ui1))
	http.Handle("/ui2/", http.StripPrefix("/ui2/", ui2))
	http.Handle("/ui3/", http.StripPrefix("/ui3/", ui3))

	http.Handle("/admin/create-user", jwtAdminAuthMiddleware(http.HandlerFunc(createUserHandler)))
	http.Handle("/admin/generate-auth-token", jwtAdminAuthMiddleware(http.HandlerFunc(generateAuthTokenHandler)))
	http.Handle("/admin/mark-user-active", jwtAdminAuthMiddleware(http.HandlerFunc(markUserActiveHandler)))
	http.Handle("/user/add-word", jwtUserAuthMiddleware(http.HandlerFunc(addWordHandler)))
	http.HandleFunc("/user/magic", keyGenHandler)

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
