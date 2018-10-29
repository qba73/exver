package main

import "testing"

func TestFixDots(t *testing.T) {
	input := "Sn.No."
	expected := "SN_NO"
	if fixDots(input) != expected {
		t.Errorf("%s != %s\n", input, expected)
	}
}

func TestFixDash(t *testing.T) {
	input := "X-Coordinates"
	expected := "X_COORDINATES"
	if fixDash(input) != expected {
		t.Errorf("%s != %s\n", input, expected)
	}
}

func TestFixSpaces(t *testing.T) {
	input := "X 23 North"
	expected := "X_23_NORTH"
	if fixSpaces(input) != expected {
		t.Errorf("%s != %s\n", input, expected)
	}
}

func TestHeaderType(t *testing.T)  {
	tp := map[string]int{
		"Sn.O": 1,
		"X-SD": 2,
		"X SD": 3,
		"XSD": 4,
	}
	for k, v := range tp {
		h := headerType(k)
		if h != v {
			t.Errorf("Incorrect header type, expected %d, got %d\n", v, h)
		}
	}

}
