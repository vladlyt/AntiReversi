package game

type Turn struct {
	Row   int
	Col   int
	SkipTurn bool
	PrintTurn bool
	Color Color
}
