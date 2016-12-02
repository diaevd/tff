package main

import (
	"strings"
	"testing"
)

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

func TestKnightLegalMoves(t *testing.T) {

	// b1
	pos := NewSquare(2, 1)
	moves := pos.KnightLegalMoves()
	str := strings.Join(moves, " ")
	if str != "Nc3 Na3 Nd2" {
		t.Errorf("KnightLegalMoves: expected [Nc3 Na3 Nd2], got [%s]", str)
	}
	// c3
	pos = NewSquare(3, 3)
	moves = pos.KnightLegalMoves()
	str = strings.Join(moves, " ")
	if str != "Nd5 Nb5 Nd1 Nb1 Ne4 Na4 Ne2 Na2" {
		t.Errorf("KnightLegalMoves: expected [Nd5 Nb5 Nd1 Nb1 Ne4 Na4 Ne2 Na2], got [%s]", str)
	}
	// a3
	pos = NewSquare(1, 3)
	moves = pos.KnightLegalMoves()
	str = strings.Join(moves, " ")
	if str != "Nb5 Nb1 Nc4 Nc2" {
		t.Errorf("KnightLegalMoves: expected [Nb5 Nb1 Nc4 Nc2], got [%s]", str)
	}
	// d2
	pos = NewSquare(4, 2)
	moves = pos.KnightLegalMoves()
	str = strings.Join(moves, " ")
	if str != "Ne4 Nc4 Nf3 Nb3 Nf1 Nb1" {
		t.Errorf("KnightLegalMoves: expected [Ne4 Nc4 Nf3 Nb3 Nf1 Nb1], got [%s]", str)
	}
}
