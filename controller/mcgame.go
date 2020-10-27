package controller

import (
	"Reversi/game"
	"fmt"
)

type MCGame struct {
	Board         game.Board
	Player        game.Color
	Enemy         game.Color
	CurrentPlayer game.Color
}

func NewMCGame(board game.Board, player game.Color) *MCGame {
	return &MCGame{
		Board:         board,
		Player:        player,
		Enemy:         !player,
		CurrentPlayer: player,
	}
}

func (g *MCGame) GetCurrentPlayer() game.Color {
	return g.CurrentPlayer
}

func (g *MCGame) GetAllTurnsOnBoard() []game.Turn {
	g.SwitchPlayers()
	return g.Board.GetAllTurns(g.CurrentPlayer)
}

func (g *MCGame) GetAllTurnsOnBoardWithSwap() []game.Turn {
	g.SwitchPlayers()
	return g.Board.GetAllTurns(g.CurrentPlayer)
}

func (g *MCGame) ExecuteTurn(turn game.Turn) {
	turn.Color = g.CurrentPlayer
	err := g.Board.UpdateBoard(turn)
	if err != nil {
		fmt.Println("ERROR in ExecuteTurn", err)
	}
	g.SwitchPlayers()
}

func (g *MCGame) SwitchPlayers() {
	g.CurrentPlayer = !g.CurrentPlayer
}

func (g *MCGame) GetWinner() game.Color {
	if g.Board.GetCellColorCount(g.Player) > g.Board.GetCellColorCount(g.Enemy) {
		return g.Player
	}
	return g.Enemy
}
