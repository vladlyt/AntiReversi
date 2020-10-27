package controller

import (
	"Reversi/game"
	"math/rand"
)

type BotPlayer struct{}

func NewBotPlayer() *BotPlayer {
	return &BotPlayer{}
}

func (p BotPlayer) Play(board game.Board, color game.Color) game.Turn {
	turns := board.GetAllTurns(color)
	if len(turns) != 0 {
		return turns[rand.Int31n(int32(len(turns)))]
	}
	return game.Turn{
		SkipTurn:  true,
		Color:     color,
		PrintTurn: true,
	}
}
