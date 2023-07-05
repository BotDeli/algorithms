package findSubstringInString

import "testing"

func TestEmpty(t *testing.T) {
	indStart := findSubstringInString("", "")
	if indStart != 0 {
		t.Errorf("Expected indStart to be 0, got %d", indStart)
	}
}

func TestEmptyStr(t *testing.T) {
	indStart := findSubstringInString("", "x")
	if indStart != -1 {
		t.Errorf("Expected indStart to be -1, got %d", indStart)
	}
}

func TestEmptySubstr(t *testing.T) {
	indStart := findSubstringInString("x", "")
	if indStart != 0 {
		t.Errorf("Expected indStart to be 0, got %d", indStart)
	}
}

func TestFullEqualsOneChar(t *testing.T) {
	indStart := findSubstringInString("x", "x")
	if indStart != 0 {
		t.Errorf("Expected indStart to be 0, got %d", indStart)
	}
}

func TestFullEqualsTwoChar(t *testing.T) {
	indStart := findSubstringInString("xx", "xx")
	if indStart != 0 {
		t.Errorf("Expected indStart to be 0, got %d", indStart)
	}
}

func TestLongSubstrTwoChar(t *testing.T) {
	indStart := findSubstringInString("x", "xx")
	if indStart != -1 {
		t.Errorf("Expected indStart to be -1, got %d", indStart)
	}
}

func TestEqualsTwoChar(t *testing.T) {
	indStart := findSubstringInString("xx", "x")
	if indStart != 0 {
		t.Errorf("Expected indStart to be 0, got %d", indStart)
	}
}

func TestSuccessfulTwoChar(t *testing.T) {
	indStart := findSubstringInString("ax", "x")
	if indStart != 1 {
		t.Errorf("Expected indStart to be 1, got %d", indStart)
	}
}

func TestFalseTwoChar(t *testing.T) {
	indStart := findSubstringInString("aa", "x")
	if indStart != -1 {
		t.Errorf("Expected indStart to be -1, got %d", indStart)
	}
}

func TestLongSubstr(t *testing.T) {
	indStart := findSubstringInString("Hello world!", "world")
	if indStart != 6 {
		t.Errorf("Expected indStart to be 6, got %d", indStart)
	}
}

func TestFindWordInStart(t *testing.T) {
	indStart := findSubstringInString("A sentence is one or more words that express a complete thought.", "A sentence")
	if indStart != 0 {
		t.Errorf("Expected indStart to be 0, got %d", indStart)
	}
}

func TestFindWord(t *testing.T) {
	indStart := findSubstringInString("A sentence is one or more words that express a complete thought.", "that")
	if indStart != 32 {
		t.Errorf("Expected indStart to be 32, got %d", indStart)
	}
}

func TestFindWordInEnd(t *testing.T) {
	indStart := findSubstringInString("A sentence is one or more words that express a complete thought.", "thought.")
	if indStart != 56 {
		t.Errorf("Expected indStart to be 56, got %d", indStart)
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

func TestLongStr(t *testing.T) {
	indStart := findSubstringInString("abeccacbadbabbad", "abbad")
	if indStart != 11 {
		t.Errorf("Expected indStart to be 11, got %d", indStart)
	}
}

func TestSingleWord(t *testing.T) {
	indStart := findSubstringInString("Metadata", "data")
	if indStart != 4 {
		t.Errorf("Expected indStart to be 4, got %d", indStart)
	}
}

func TestUpperSubstr(t *testing.T) {
	indStart := findSubstringInString("Metadata", "DATA")
	if indStart != -1 {
		t.Errorf("Expected indStart to be -1, got %d", indStart)
	}
}

func TestLowerSubstr(t *testing.T) {
	indStart := findSubstringInString("MetaData", "data")
	if indStart != -1 {
		t.Errorf("Expected indStart to be -1, got %d", indStart)
	}
}

func TestUpperStr(t *testing.T) {
	indStart := findSubstringInString("METADATA", "data")
	if indStart != -1 {
		t.Errorf("Expected indStart to be -1, got %d", indStart)
	}
}

func TestSuccessfulUpperStr(t *testing.T) {
	indStart := findSubstringInString("METADATA", "DATA")
	if indStart != 4 {
		t.Errorf("Expected indStart to be 4, got %d", indStart)
	}
}
