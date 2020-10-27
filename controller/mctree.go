package controller

import (
	"Reversi/game"
	"math"
)

type MCTree struct {
	parent        *MCTree
	children      []*MCTree
	turn          game.Turn
	wins          int
	visitedCount  int
	currentPlayer game.Color
	freeTurns     []game.Turn
}

func NewMCTree(parent *MCTree, turn game.Turn, game *MCGame) *MCTree {
	return &MCTree{
		children:      make([]*MCTree, 0),
		parent:        parent,
		turn:          turn,
		freeTurns:     game.Board.GetAllTurns(game.GetCurrentPlayer()),
		currentPlayer: game.GetCurrentPlayer(),
	}
}

func (t *MCTree) MostVisitedTurn() *game.Turn {
	if len(t.children) == 0 {
		return nil
	}
	mostVisitedChildValue := -1
	mostVisitedChildIdx := 0
	for i := range t.children {
		if t.children[i].visitedCount > mostVisitedChildValue {
			mostVisitedChildIdx = i
			mostVisitedChildValue = t.children[i].visitedCount
		}
	}
	return &t.children[mostVisitedChildIdx].turn
}

func (t *MCTree) GetNextNode() *MCTree {
	var node *MCTree = nil
	best := math.Inf(-1)
	for _, child := range t.children {
		uct := t.GetUCT(child)
		if uct > best {
			node = child
			best = uct
		}
	}
	return node
}

func (t *MCTree) AddChild(g *MCGame, index int) *MCTree {
	node := NewMCTree(t, t.freeTurns[index], g)
	copy(t.freeTurns[index:], t.freeTurns[index+1:])
	t.freeTurns[len(t.freeTurns)-1] = game.Turn{}
	t.freeTurns = t.freeTurns[:len(t.freeTurns)-1]
	t.children = append(t.children, node)
	return node
}

func (t *MCTree) UpdateScoreInNode(winner game.Color) {
	t.visitedCount++
	if winner == t.currentPlayer {
		t.wins++
	}
}

func (t *MCTree) GetUCT(child *MCTree) float64 {
	return float64(child.wins)/float64(child.visitedCount) + math.Sqrt(2*math.Log(float64(t.visitedCount))/float64(child.visitedCount))
}
