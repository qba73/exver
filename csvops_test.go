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
