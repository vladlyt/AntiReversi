package view

import (
	"Reversi/game"
	"fmt"
	"strconv"
	"strings"
)

type GameView struct {
}

type View interface {
	Run(events <-chan game.Event)
}

func NewGameView() *GameView {
	return &GameView{

	}
}

func OutputFromTurn(turn game.Turn) string {
	if turn.SkipTurn {
		return "pass"
	}
	return strings.ToUpper(string(rune(97+turn.Col))) + strconv.Itoa(turn.Row+1)
}

func (view *GameView) Run(events <-chan game.Event) {
	for event := range events {
		switch event.Event {
		case game.NewTurn:
			fmt.Println(OutputFromTurn(event.Turn))
		default:
		}
	}
}
