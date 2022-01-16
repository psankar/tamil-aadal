package main

import (
	"fmt"
	"strings"
)

func isUyir(c string) bool {
	return strings.Contains("அஆஇஈஉஊஎஏஐஒஓஔ", c)
}

func isAayudham(c string) bool {
	return c == "ஃ"
}

func getMei(c string) string {
	var meiMap map[string]string = map[string]string{
		"க": "க்",
		"ங": "ங்",
		"ச": "ச்",
		"ஞ": "ஞ்",
		"ட": "ட்",
		"ண": "ண்",
		"த": "த்",
		"ந": "ந்",
		"ப": "ப்",
		"ம": "ம்",
		"ய": "ய்",
		"ர": "ர்",
		"ல": "ல்",
		"வ": "வ்",
		"ழ": "ழ்",
		"ள": "ள்",
		"ற": "ற்",
		"ன": "ன்",
	}

	return meiMap[string([]rune(c)[0])]
}

func getUyir(c string) string {
	var uyirMap map[string]string = map[string]string{
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

	var diacriticMap map[string]string = map[string]string{
		string('\u0BBE'): "ஆ",
		string('\u0BBF'): "இ",
		string('\u0BC0'): "ஈ",
		string('\u0BC1'): "உ",
		string('\u0BC2'): "ஊ",
		string('\u0BC6'): "எ",
		string('\u0BC7'): "ஏ",
		string('\u0BC8'): "ஐ",
		string('\u0BCA'): "ஒ",
		string('\u0BCB'): "ஓ",
		string('\u0BCC'): "ஔ",
	}

	if uyirMap[c] == "" {
		var diacritic = []rune(c)
		if len(diacritic) > 1 && !isAayudham(string(diacritic[0])) && !isUyir(string(diacritic[0])) && getMei(string(diacritic[0])) != "" {
			return diacriticMap[string(diacritic[1])]
		}
	}
	return uyirMap[c]
}

func validate(gotLetters []string, wantLetters []string) []string {
	var results []string

	if len(gotLetters) != len(wantLetters) {
		fmt.Println("Please make wantLetters and gotLetters of equal length")
		return results
	}

nextLetter:
	for i, letter := range gotLetters {

		// complete match
		if letter == wantLetters[i] {
			results = append(results, LetterMatched)
			continue
		}

		// Found else where
		for _, j := range wantLetters {
			if j == letter {
				results = append(results, LetterElseWhere)
				continue nextLetter
			}
		}

		// Mei match
		if getMei(letter) == getMei(wantLetters[i]) && getMei(letter) != "" {
			results = append(results, MeiMatched)
			continue nextLetter
		}

		// uyir match
		if getUyir(letter) == getUyir(wantLetters[i]) && getUyir(letter) != "" {
			results = append(results, UyirMatched)
			continue nextLetter
		}

		// no match
		results = append(results, LetterNotFound)
		continue nextLetter
	}

	return results
}
