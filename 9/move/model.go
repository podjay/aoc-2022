package move

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type It struct {
	Direction Direction
	Distance  int
}

type Position struct {
	X int
	Y int
}
