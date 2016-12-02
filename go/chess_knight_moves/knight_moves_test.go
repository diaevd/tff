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

	// b1
	pos := NewSquare(2, 1)
	moves := pos.KnightLegalMoves()
	str := strings.Join(moves, " ")
	res := "Nc3 Na3 Nd2"
	if str != res {
		t.Errorf("KnightLegalMoves: expected [%s], got [%s]", res, str)
	}
	// c3
	pos = NewSquare(3, 3)
	moves = pos.KnightLegalMoves()
	str = strings.Join(moves, " ")
	res = "Nd5 Nb5 Nd1 Nb1 Ne4 Na4 Ne2 Na2"
	if str != res {
		t.Errorf("KnightLegalMoves: expected [%s], got [%s]", res, str)
	}
	// a3
	pos = NewSquare(1, 3)
	moves = pos.KnightLegalMoves()
	str = strings.Join(moves, " ")
	res = "Nb5 Nb1 Nc4 Nc2"
	if str != res {
		t.Errorf("KnightLegalMoves: expected [%s], got [%s]", res, str)
	}
	// d2
	pos = NewSquare(4, 2)
	moves = pos.KnightLegalMoves()
	str = strings.Join(moves, " ")
	res = "Ne4 Nc4 Nf3 Nb3 Nf1 Nb1"
	if str != res {
		t.Errorf("KnightLegalMoves: expected [%s], got [%s]", res, str)
	}
}
