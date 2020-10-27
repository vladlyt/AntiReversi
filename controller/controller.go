package controller

import (
	"Reversi/game"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Players map[game.Color]Player

type GameRunner struct {
}

type Controller interface {
	Run(boards chan<- game.Event)
}

func NewGameRunner() *GameRunner {
	return &GameRunner{}
}

func TurnFromInput(input string) game.Turn {
	if strings.HasPrefix(input, "pass") {
		return game.Turn{
			SkipTurn: true,
		}
	}
	row, _ := strconv.Atoi(input[1:2])
	col := int(strings.ToLower(input)[0]) - 97
	return game.Turn{
		Row: row - 1,
		Col: col,
	}
}

func StringColorToColor(color string) game.Color {
	return game.Color(strings.HasPrefix(color, "white"))
}

func (runner *GameRunner) initPlayers(color game.Color, reader *bufio.Reader) Players {
	players := make(Players)

	players[color] = NewMonteCarloBot(time.Second, color)
	players[!color] = NewRealPlayer(reader)

	return players
}

func (runner *GameRunner) Run(events chan<- game.Event) {
	defer close(events)

	reader := bufio.NewReader(os.Stdin)
	blackHole, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	color, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	players := runner.initPlayers(StringColorToColor(color), reader)
	gameModel := game.NewGame(events)
	gameModel.SetBlackHole(TurnFromInput(blackHole))

	for !gameModel.IsGameFinished() {
		currentColor := gameModel.CurrentColor
		turn := players[currentColor].Play(gameModel.Board, currentColor)

		err := gameModel.DoTurn(turn)
		if err != nil {
			fmt.Println("ERROR in run", err)
		}
	}
	data, _ := reader.ReadString('\n')
	for data != "" {
		data, _ = reader.ReadString('\n')
	}
}
