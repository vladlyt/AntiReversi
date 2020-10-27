package controller

import (
	"Reversi/game"
	"bufio"
)

type RealPlayer struct {
	reader *bufio.Reader
}

func NewRealPlayer(reader *bufio.Reader) *RealPlayer {
	return &RealPlayer{
		reader: reader,
	}
}

func (p RealPlayer) Play(board game.Board, color game.Color) game.Turn {
	input, err := p.reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	turn := TurnFromInput(input)
	turn.Color = color
	return turn
}
