package game

const (
	NewTurn = iota
	BoardUpdate
	WinnerScreen
)

type Event struct {
	Event       int
	Board       Board
	Turn        Turn
	WhiteResult int
	BlackResult int
}
