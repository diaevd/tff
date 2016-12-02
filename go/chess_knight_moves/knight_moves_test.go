package main

import (
	"strings"
	"testing"
)

func TestSquareString(t *testing.T) {

	s1 := Square{1, 1}
	str1 := s1.String()
	res1 := "a1"
	if str1 != res1 {
		t.Errorf("SquareString: expected %s, got %s", res1, str1)
	}
	s2 := Square{8, 8}
	str2 := s2.String()
	res2 := "h8"
	if str2 != res2 {
		t.Errorf("SquareString: expected %s, got %s", res2, str2)
	}
}

func TestKnightLegalMoves(t *testing.T) {

	places := []Square{
		{2, 1}, // b1
		{3, 3}, // c3
		{1, 3}, // a3
		{4, 2}, // d2
	}
	expected := []string{
		"Nc3 Na3 Nd2",                     // b1
		"Nd5 Nb5 Nd1 Nb1 Ne4 Na4 Ne2 Na2", // c3
		"Nb5 Nb1 Nc4 Nc2",                 // a3
		"Ne4 Nc4 Nf3 Nb3 Nf1 Nb1",         // d2
	}
	for i, pos := range places {
		// b1
		moves := pos.KnightLegalMoves()
		str := strings.Join(moves, " ")
		if str != expected[i] {
			t.Errorf("KnightLegalMoves: expected [%s], got [%s]", expected[i], str)
		}
	}
}
