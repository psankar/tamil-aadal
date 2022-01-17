package main

import (
	"reflect"
	"testing"
)

var basicTests = []struct {
	input      string
	isUyir     bool
	isAayudham bool
	mei        string
	uyir       string
}{
	{"அ", true, false, "", "அ"},
	{"ஆ", true, false, "", "ஆ"},
	{"இ", true, false, "", "இ"},
	{"ஈ", true, false, "", "ஈ"},
	{"உ", true, false, "", "உ"},
	{"ஊ", true, false, "", "ஊ"},
	{"எ", true, false, "", "எ"},
	{"ஏ", true, false, "", "ஏ"},
	{"ஐ", true, false, "", "ஐ"},
	{"ஒ", true, false, "", "ஒ"},
	{"ஓ", true, false, "", "ஓ"},
	{"ஔ", true, false, "", "ஔ"},
	{"ஃ", false, true, "", ""},
	{"க", false, false, "க்", "அ"},
	{"ங", false, false, "ங்", "அ"},
	{"ச", false, false, "ச்", "அ"},
	{"ஞ", false, false, "ஞ்", "அ"},
	{"ட", false, false, "ட்", "அ"},
	{"ண", false, false, "ண்", "அ"},
	{"த", false, false, "த்", "அ"},
	{"ந", false, false, "ந்", "அ"},
	{"ப", false, false, "ப்", "அ"},
	{"ம", false, false, "ம்", "அ"},
	{"ய", false, false, "ய்", "அ"},
	{"ர", false, false, "ர்", "அ"},
	{"ல", false, false, "ல்", "அ"},
	{"வ", false, false, "வ்", "அ"},
	{"ழ", false, false, "ழ்", "அ"},
	{"ள", false, false, "ள்", "அ"},
	{"ற", false, false, "ற்", "அ"},
	{"ன", false, false, "ன்", "அ"},
	{"க்", false, false, "க்", ""},
	{"ங்", false, false, "ங்", ""},
	{"ச்", false, false, "ச்", ""},
	{"ஞ்", false, false, "ஞ்", ""},
	{"ட்", false, false, "ட்", ""},
	{"ண்", false, false, "ண்", ""},
	{"த்", false, false, "த்", ""},
	{"ந்", false, false, "ந்", ""},
	{"ப்", false, false, "ப்", ""},
	{"ம்", false, false, "ம்", ""},
	{"ய்", false, false, "ய்", ""},
	{"ர்", false, false, "ர்", ""},
	{"ல்", false, false, "ல்", ""},
	{"வ்", false, false, "வ்", ""},
	{"ழ்", false, false, "ழ்", ""},
	{"ள்", false, false, "ள்", ""},
	{"ற்", false, false, "ற்", ""},
	{"ன்", false, false, "ன்", ""},
	{"கா", false, false, "க்", "ஆ"},
	{"கி", false, false, "க்", "இ"},
	{"கீ", false, false, "க்", "ஈ"},
	{"கு", false, false, "க்", "உ"},
	{"கூ", false, false, "க்", "ஊ"},
	{"கெ", false, false, "க்", "எ"},
	{"கே", false, false, "க்", "ஏ"},
	{"கை", false, false, "க்", "ஐ"},
	{"கொ", false, false, "க்", "ஒ"},
	{"கோ", false, false, "க்", "ஓ"},
	{"கௌ", false, false, "க்", "ஔ"},
	{"அ்", false, false, "", ""},
	{"அா", false, false, "", ""},
	{"அி", false, false, "", ""},
	{"அீ", false, false, "", ""},
	{"அு", false, false, "", ""},
	{"அூ", false, false, "", ""},
	{"அெ", false, false, "", ""},
	{"அே", false, false, "", ""},
	{"அை", false, false, "", ""},
	{"அொ", false, false, "", ""},
	{"அோ", false, false, "", ""},
	{"அௌ", false, false, "", ""},
	{"ஃ்", false, false, "", ""},
	{"ஃா", false, false, "", ""},
	{"ஃி", false, false, "", ""},
	{"ஃீ", false, false, "", ""},
	{"ஃு", false, false, "", ""},
	{"ஃூ", false, false, "", ""},
	{"ஃெ", false, false, "", ""},
	{"ஃே", false, false, "", ""},
	{"ஃை", false, false, "", ""},
	{"ஃொ", false, false, "", ""},
	{"ஃோ", false, false, "", ""},
	{"ஃௌ", false, false, "", ""},
	{"்்", false, false, "", ""},
	{"்ா", false, false, "", ""},
	{"்ி", false, false, "", ""},
	{"்ீ", false, false, "", ""},
	{"்ு", false, false, "", ""},
	{"்ூ", false, false, "", ""},
	{"்ெ", false, false, "", ""},
	{"்ே", false, false, "", ""},
	{"்ை", false, false, "", ""},
	{"்ொ", false, false, "", ""},
	{"்ோ", false, false, "", ""},
	{"்ௌ", false, false, "", ""},
	{"ா", false, false, "", ""},
	{"ி", false, false, "", ""},
	{"ீ", false, false, "", ""},
	{"ு", false, false, "", ""},
	{"ூ", false, false, "", ""},
	{"ெ", false, false, "", ""},
	{"ே", false, false, "", ""},
	{"ை", false, false, "", ""},
	{"ொ", false, false, "", ""},
	{"ோ", false, false, "", ""},
	{"ௌ", false, false, "", ""},
	{"ிீ", false, false, "", ""},
	{"a", false, false, "", ""},
}

func TestGetMei(t *testing.T) {
	for _, test := range basicTests {
		if got := meiMap[test.input]; got != test.mei {
			t.Errorf("meiMap[%q] = %v, want %v", test.input, got, test.mei)
		}
	}
}

func TestGetUyir(t *testing.T) {
	for _, test := range basicTests {
		if got := uyirMap[test.input]; got != test.uyir {
			t.Errorf("uyirMap[%q] = %v, want %v", test.input, got, test.uyir)
		}
	}
}

var wordTests = []struct {
	left       []string
	right      []string
	result     [][]string
	allMatched bool
}{
	{
		[]string{"ப", "ல்", "ம", "ணி", "மி", "கா", "லை", "ஆ", "ஃ", "ஃ"},
		[]string{"அ", "ன்", "ம", "ணி", "மே", "க", "லை", "ஆ", "டு", "ஃ"},
		[][]string{{LetterNotFound, UyirMatched}, {LetterNotFound}, {LetterMatched}, {LetterMatched}, {LetterNotFound, MeiMatched}, {LetterNotFound, MeiMatched}, {LetterMatched}, {LetterMatched}, {LetterElseWhere}, {LetterMatched}},
		false,
	},
	{
		[]string{"ப", "ல்", "ம", "ணி", "மி", "கா", "லை", "ஆ", "ஃ", "a"},
		[]string{"அ", "ன்", "ம", "ணி", "மே", "க", "லை", "ஆ", "டு", "ஃ"},
		[][]string{{LetterNotFound, UyirMatched}, {LetterNotFound}, {LetterMatched}, {LetterMatched}, {LetterNotFound, MeiMatched}, {LetterNotFound, MeiMatched}, {LetterMatched}, {LetterMatched}, {LetterElseWhere}, {LetterNotFound}},
		false,
	},
	{
		[]string{"த", "மி", "ழ்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterMatched}, {LetterMatched}},
		true,
	},
	{
		[]string{"த", "மி", "ழ"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterMatched}, {LetterNotFound, MeiMatched}},
		false,
	},
	{
		[]string{"த", "மி", "ல்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterMatched}, {LetterNotFound}},
		false,
	},
	{
		[]string{"த", "மி", "ழு"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterMatched}, {LetterNotFound, MeiMatched}},
		false,
	},
	{
		[]string{"த", "மா", "ழ"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterNotFound, MeiMatched}, {LetterNotFound, MeiMatched}},
		false,
	},
	{
		[]string{"த", "ரி", "ழ்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterNotFound, UyirMatched}, {LetterMatched}},
		false,
	},
	{
		[]string{"அ", "ரி", "ழ்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterNotFound, UyirMatched}, {LetterNotFound, UyirMatched}, {LetterMatched}},
		false,
	},
	{
		[]string{"ச", "ரி", "ழ்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterNotFound, UyirMatched}, {LetterNotFound, UyirMatched}, {LetterMatched}},
		false,
	},
	{
		[]string{"ச", "மி", "ழ்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterNotFound, UyirMatched}, {LetterMatched}, {LetterMatched}},
		false,
	},
	{
		[]string{"தா", "மி", "ழ்"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterNotFound, MeiMatched}, {LetterMatched}, {LetterMatched}},
		false,
	},
	{
		[]string{"த", "ழ்", "மி"},
		[]string{"த", "மி", "ழ்"},
		[][]string{{LetterMatched}, {LetterElseWhere}, {LetterElseWhere}},
		false,
	},
	{
		[]string{"த", "ழி", "ழி", "ழ்"},
		[]string{"த", "மி", "ழி", "ல்"},
		[][]string{{LetterMatched}, {LetterElseWhere, UyirMatched}, {LetterMatched}, {LetterNotFound}},
		false,
	},
}

func TestVerifyWordWithUyirMei(t *testing.T) {
	var empty struct{}
	for _, test := range wordTests {

		// TODO: This is a hack to warm up the map. `verifyWordWithUyirMei` should not depend on the local map.
		todayLettersMap = make(map[string]struct{})
		for _, letter := range test.right {
			todayLettersMap[letter] = empty
		}

		gotResults, gotAllMatched := verifyWordWithUyirMei(test.left, test.right)
		if !reflect.DeepEqual(gotResults, test.result) {
			t.Errorf("verifyWordWithUyirMei(%q, %q) = %v, want %v", test.left, test.right, gotResults, test.result)
		}
		if gotAllMatched != test.allMatched {
			t.Errorf("verifyWordWithUyirMei(%q, %q) = %v, want %v", test.left, test.right, gotAllMatched, test.allMatched)
		}
	}
}
