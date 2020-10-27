package game

type Game struct {
	CurrentColor Color
	Board        Board
	events       chan<- Event
}

func NewGame(events chan<- Event) *Game {
	board := NewBoard()
	return &Game{
		CurrentColor: BLACK,
		Board:        board,
		events:       events,
	}
}

func (g *Game) SetBlackHole(turn Turn) {
	g.Board[turn.Row][turn.Col].IsBlackHole = true
}

func (g *Game) IsGameFinished() bool {
	return !g.HasTurns(BLACK) && !g.HasTurns(WHITE)
}

func (g *Game) HasTurns(color Color) bool {
	return g.Board.HasTurns(color)
}

func (g *Game) SetNextPlayer() {
	g.CurrentColor = !g.CurrentColor
}

func (g *Game) GetResults() (white int, black int) {
	return g.Board.GetCellColorCount(WHITE), g.Board.GetCellColorCount(BLACK)
}

func (g *Game) DoTurn(turn Turn) error {
	err := g.Board.UpdateBoard(turn)
	if err != nil {
		return err
	}
	if turn.PrintTurn {
		g.events <- Event{
			Event: NewTurn,
			Turn:  turn,
		}
	}

	//fmt.Println(g.Board)
	g.SetNextPlayer()
	return nil
}
