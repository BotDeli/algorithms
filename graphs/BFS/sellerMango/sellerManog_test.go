package sellerMango

import (
	"testing"
)

func TestNil(t *testing.T) {
	data := &Seller{
		Sell:      false,
		Neighbors: nil,
		Level:     0,
	}
	level := FindSellerMango(data)
	expected := -1
	if level != expected {
		t.Errorf("Expected %d, got %d", expected, level)
	}
}

func TestISeller(t *testing.T) {
	data := &Seller{
		Sell:      true,
		Neighbors: nil,
		Level:     0,
	}
	level := FindSellerMango(data)
	expected := 0
	if level != expected {
		t.Errorf("Expected %d, got %d", expected, level)
	}
}

func TestFalseSeller(t *testing.T) {
	data := &Seller{
		Sell: false,
		Neighbors: []*Seller{
			{false, nil, 1},
			{false, nil, 1},
			{false, nil, 1},
			{false, nil, 1},
			{false, []*Seller{
				{false, nil, 2},
				{false, nil, 2},
			}, 1},
		},
		Level: 0,
	}
	level := FindSellerMango(data)
	expected := -1
	if level != expected {
		t.Errorf("Expected %d, got %d", expected, level)
	}
}

func TestTrueSellerOneLevel(t *testing.T) {
	data := &Seller{
		Sell: false,
		Neighbors: []*Seller{
			{false, nil, 1},
			{true, nil, 1},
			{false, nil, 1},
			{false, nil, 1},
			{false, []*Seller{
				{false, nil, 2},
				{false, nil, 2},
			}, 1},
		},
		Level: 0,
	}
	level := FindSellerMango(data)
	expected := 1
	if level != expected {
		t.Errorf("Expected %d, got %d", expected, level)
	}
}

func TestTrueSellerTwoLevel(t *testing.T) {
	data := &Seller{
		Sell: false,
		Neighbors: []*Seller{
			{false, nil, 1},
			{false, []*Seller{
				{false, nil, 2},
				{true, nil, 2},
			}, 1},
		},
		Level: 0,
	}
	level := FindSellerMango(data)
	expected := 2
	if level != expected {
		t.Errorf("Expected %d, got %d", expected, level)
	}
}

func TestTrueSellerThreeLevel(t *testing.T) {
	data := &Seller{
		Sell: false,
		Neighbors: []*Seller{
			{false, nil, 1},
			{false, []*Seller{
				{false, nil, 1},
				{false, []*Seller{
					{false, nil, 3},
					{true, nil, 3},
				}, 2},
			}, 1},
		},
		Level: 0,
	}
	level := FindSellerMango(data)
	expected := 3
	if level != expected {
		t.Errorf("Expected %d, got %d", expected, level)
	}
}
