package bankCard

import "testing"

func TestFalseCountNumber(t *testing.T) {
	card := "1234 5678 9012 345"
	if isValidCardNumber(card) {
		t.Fail()
	}
}

func TestTrueCheckNumber(t *testing.T) {
	card := "4893 4703 1885 1588"
	if !isValidCardNumber(card) {
		t.Fail()
	}
}

func TestFalseCheckNumber(t *testing.T) {
	card := "4993 4703 1885 1588"
	if isValidCardNumber(card) {
		t.Fail()
	}
}
