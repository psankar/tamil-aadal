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
