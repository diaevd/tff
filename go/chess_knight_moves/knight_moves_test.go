package main

import "testing"

func TestSquareString(t *testing.T) {

	s1 := Square{1, 1}
	str1 := s1.String()
	if str1 != "a1" {
		t.Errorf("SquareString: expected a1, got %s", str1)
	}
	s2 := Square{8, 8}
	str2 := s2.String()
	if str2 != "h8" {
		t.Errorf("SquareString: expected a1, got %s", str2)
	}
}
