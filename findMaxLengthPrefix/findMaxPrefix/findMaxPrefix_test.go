package findMaxPrefix

import "testing"

func TestOneChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("a")
	if maxLength != 0 {
		t.Errorf("expected maxLength to be 0, got %d", maxLength)
	}
}

func TestDontEqualsTwoChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("ab")
	if maxLength != 0 {
		t.Errorf("expected maxLength to be 0, got %d", maxLength)
	}
}

func TestEqualsTwoChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("aa")
	if maxLength != 1 {
		t.Errorf("expected maxLength to be 1, got %d", maxLength)
	}
}

func TestThreeChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("aba")
	if maxLength != 1 {
		t.Errorf("expected maxLength to be 1, got %d", maxLength)
	}
}

func TestFourChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("abab")
	if maxLength != 2 {
		t.Errorf("expected maxLength to be 2, got %d", maxLength)
	}
}

func TestFiveChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("ababa")
	if maxLength != 3 {
		t.Errorf("expected maxLength to be 3, got %d", maxLength)
	}
}

func TestNineChar(t *testing.T) {
	maxLength := findMaxLengthPrefix("abbaabba")
	if maxLength != 4 {
		t.Errorf("expected maxLength to be 4, got %d", maxLength)
	}
}

func TestLongStr(t *testing.T) {
	maxLength := findMaxLengthPrefix("abccdeujbkabccdeuj")
	if maxLength != 8 {
		t.Errorf("expected maxLength to be 8, got %d", maxLength)
	}
}

func TestFindMaxSuffix(t *testing.T) {
	maxLength := findMaxLengthPrefix("lilila")
	if maxLength != 3 {
		t.Errorf("expected maxLength to be 4, got %d", maxLength)
	}
}
