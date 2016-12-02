package main

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

func main() {
}
