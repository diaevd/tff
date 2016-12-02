package main

import "log"

// Square x, y coordinates in board
type Square struct {
	X, Y int
}

var (
	// Vert вертикаль
	Vert = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	// Horiz горизонталь
	Horiz = []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
)

// String as is
func (s *Square) String() string {
	bytearray := [2]byte{Vert[s.X-1], Horiz[s.Y-1]}
	return string(bytearray[:])
}

// NewSquare конструктор новой карты
func NewSquare(x int, y int) *Square {

	if x < 1 || x > 8 {
		panic("X is out of range")
	}
	if y < 1 || y > 8 {
		panic("Y is out of range")
	}

	return &Square{X: x, Y: y}
}

// Board доска
type Board []*Square

// NewBoard конструктор доски
// Вдруг решим доску все-таки заполнить, пригодится
func NewBoard() Board {

	board := Board{}

	for i := range Vert {
		for j := range Horiz {
			board = append(board, NewSquare(i+1, j+1))
		}
	}

	return board
}

// KnightLegalMoves ...
func (s Square) KnightLegalMoves() []string {

	// возможные направления движения
	directions := [][2]int{
		{1, 2},
		{-1, 2},
		{1, -2},
		{-1, -2},
		{2, 1},
		{-2, 1},
		{2, -1},
		{-2, -1},
	}

	moves := []string{}
	for _, d := range directions {
		x := s.X + d[0]
		y := s.Y + d[1]
		// check for out of ranges
		if x < 1 || x > 8 {
			continue
		}
		if y < 1 || y > 8 {
			continue
		}
		moves = append(moves, "N"+NewSquare(x, y).String())
	}

	return moves
}

func main() {

	// b1
	pos := NewSquare(2, 1)
	moves := pos.KnightLegalMoves()
	log.Println("pos:", pos, " moves:", moves)
	// c3
	pos = NewSquare(3, 3)
	moves = pos.KnightLegalMoves()
	log.Println("pos:", pos, " moves:", moves)
	// a3
	pos = NewSquare(1, 3)
	moves = pos.KnightLegalMoves()
	log.Println("pos:", pos, " moves:", moves)
	// d2
	pos = NewSquare(4, 2)
	moves = pos.KnightLegalMoves()
	log.Println("pos:", pos, " moves:", moves)
}
