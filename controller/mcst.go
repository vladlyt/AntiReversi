package controller

import (
	"Reversi/game"
	"math/rand"
	"time"
)

type MonteCarloBot struct {
	maxTime        time.Duration
	node           *MCTree
	gameSimulation *MCGame
	color          game.Color
}

func NewMonteCarloBot(maxTime time.Duration, color game.Color) *MonteCarloBot {
	return &MonteCarloBot{
		maxTime: maxTime,
		color:   color,
	}
}

func (p MonteCarloBot) Play(board game.Board, color game.Color) game.Turn {
	root := NewMCTree(nil, game.Turn{}, NewMCGame(board.Copy(), color))
	deadline := time.Now().Add(p.maxTime)

	for time.Until(deadline) > 0 {
		p.node = root
		p.gameSimulation = NewMCGame(board.Copy(), color)
		p.selection()
		p.expansion()
		p.simulation()
		p.backPropagation()
	}

	mvc := root.MostVisitedTurn()
	if mvc == nil {
		return game.Turn{
			SkipTurn:  true,
			Color:     color,
			PrintTurn: true,
		}
	}
	return *mvc
}

func (p *MonteCarloBot) selection() {
	for len(p.node.freeTurns) == 0 && len(p.node.children) > 0 {
		p.node = p.node.GetNextNode()
		p.gameSimulation.ExecuteTurn(p.node.turn)
	}
}

func (p *MonteCarloBot) expansion() {
	for len(p.node.freeTurns) > 0 {
		idx := rand.Int() % len(p.node.freeTurns)
		p.gameSimulation.ExecuteTurn(p.node.freeTurns[idx])
		p.node = p.node.AddChild(p.gameSimulation, idx)
	}

}
func (p *MonteCarloBot) simulation() {
	availableTurns := p.gameSimulation.GetAllTurnsOnBoard()
	for len(availableTurns) > 0 {
		p.gameSimulation.ExecuteTurn(availableTurns[rand.Int()%len(availableTurns)])
		availableTurns = p.gameSimulation.GetAllTurnsOnBoard()
		if len(availableTurns) == 0 {
			availableTurns = p.gameSimulation.GetAllTurnsOnBoardWithSwap()
		}
	}
}

func (p *MonteCarloBot) backPropagation() {
	winner := p.gameSimulation.GetWinner()
	for p.node != nil {
		p.node.UpdateScoreInNode(winner)
		p.node = p.node.parent
	}
}
