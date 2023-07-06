package Knuta

import (
	"testing"
)

const (
	ZERO       = 0
	FalseFound = -1
)

func TestEmpty(t *testing.T) {
	indStart := findSubstringInString("", "")
	if indStart != ZERO {
		t.Errorf("Expected indStart to be %d, got %d", ZERO, indStart)
	}
}

func TestEmptyStr(t *testing.T) {
	indStart := findSubstringInString("", " ")
	if indStart != FalseFound {
		t.Errorf("Expected indStart to be %d, got %d", FalseFound, indStart)
	}
}

func TestEmptySubstr(t *testing.T) {
	indStart := findSubstringInString(" ", "")
	if indStart != ZERO {
		t.Errorf("Expected indStart to be %d, got %d", ZERO, indStart)
	}
}

func TestFullEqualsOneChar(t *testing.T) {
	indStart := findSubstringInString("a", "a")
	if indStart != ZERO {
		t.Errorf("Expected indStart to be %d, got %d", ZERO, indStart)
	}
}

func TestFullEqualsTwoChars(t *testing.T) {
	indStart := findSubstringInString("ab", "ab")
	if indStart != ZERO {
		t.Errorf("Expected indStart to be %d, got %d", ZERO, indStart)
	}
}

func TestFullEqualsThreeChars(t *testing.T) {
	indStart := findSubstringInString("abc", "abc")
	if indStart != ZERO {
		t.Errorf("Expected indStart to be %d, got %d", ZERO, indStart)
	}
}

func TestTrueFoundOneChar(t *testing.T) {
	indStart := findSubstringInString("ab", "b")
	expectedIndStart := 1
	if indStart != expectedIndStart {
		t.Errorf("Expected indStart to be %d, got %d", expectedIndStart, indStart)
	}
}

func TestFalseFoundOneChar(t *testing.T) {
	indStart := findSubstringInString("ab", "c")
	if indStart != FalseFound {
		t.Errorf("Expected indStart to be %d, got %d", FalseFound, indStart)
	}
}

func TestTrueFoundTwoChars(t *testing.T) {
	indStart := findSubstringInString("cdab", "ab")
	expectedIndStart := 2
	if indStart != expectedIndStart {
		t.Errorf("Expected indStart to be %d, got %d", expectedIndStart, indStart)
	}
}

func TestFalseFoundTwoChars(t *testing.T) {
	indStart := findSubstringInString("dac", "cd")
	if indStart != FalseFound {
		t.Errorf("Expected indStart to be %d, got %d", FalseFound, indStart)
	}
}

func TestTrueFoundLongStr(t *testing.T) {
	indStart := findSubstringInString("abcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabbcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd", "bb")
	expectedIndStart := 293
	if indStart != expectedIndStart {
		t.Errorf("Expected indStart to be %d, got %d", expectedIndStart, indStart)
	}
}

func TestRepeatedWord(t *testing.T) {
	indStart := findSubstringInString("A sentence is one or more WordsWordsWords that express a complete thought.", "Words")
	if indStart != 26 {
		t.Errorf("Expected indStart to be 26, got %d", indStart)
	}
}

func TestGrowWords(t *testing.T) {
	indStart := findSubstringInString("W Wo Wor Word Words", "Words")
	if indStart != 14 {
		t.Errorf("Expected indStart to be 14, got %d", indStart)
	}
}
