package controller

import (
	"Reversi/game"
	"fmt"
	"os"
)

const MAX_SCORE = 1000

type MinimaxBot struct {
	depth  int
	cache  map[game.Board]int
	scores map[[2]int]int
	f      *os.File
}

func NewMinimaxBotPlayer(depth int, f *os.File) *MinimaxBot {
	return &MinimaxBot{
		depth: depth,
		f:     f,
		cache: make(map[game.Board]int),
		scores: map[[2]int]int{
			{0, 0}: -99,
			{0, 1}: 48,
			{0, 2}: -8,
			{0, 3}: 6,
			{1, 0}: 48,
			{1, 1}: -8,
			{1, 2}: -16,
			{1, 3}: 3,
			{2, 0}: -8,
			{2, 1}: -16,
			{2, 2}: 4,
			{2, 3}: 4,
			{3, 0}: 6,
			{3, 1}: 3,
			{3, 2}: 4,
			{3, 3}: 0,
		},
	}
}

func (p MinimaxBot) Play(board game.Board, color game.Color) game.Turn {
	turns := board.GetAllTurns(color)

	if len(turns) == 0 {
		return game.Turn{
			SkipTurn:  true,
			Color:     color,
			PrintTurn: true,
		}
	}

	bestTurn := turns[0]
	heurScore := -MAX_SCORE
	for _, turn := range turns {
		newBoard := board.Copy()
		err := newBoard.UpdateBoard(turn)
		if err != nil {
			fmt.Println("ERROR IN MINIMAX TURN", err)
		}
		score := p.minGameplay(newBoard, color, p.depth-1)
		if score > heurScore {
			bestTurn = turn
			heurScore = score
		}

	}
	return bestTurn
}

func (p *MinimaxBot) calculateHeuristics(board game.Board, color game.Color, depth int) int {
	score := 0
	tmpIndexes := [2]int{0, 0}
	for row := range board {
		for col := range board[row] {
			if board[row][col].IsFilled && !board[row][col].IsBlackHole {
				if row < 4 && col < 4 {
					tmpIndexes[0] = row
					tmpIndexes[1] = col
				} else if row >= 4 && col < 4 {
					tmpIndexes[0] = row - 4
					tmpIndexes[1] = col
				} else if row < 4 && col >= 4 {
					tmpIndexes[0] = row
					tmpIndexes[1] = col - 4
				} else {
					tmpIndexes[0] = row - 4
					tmpIndexes[1] = col - 4
				}
				if board[row][col].Color == color {
					score += p.scores[tmpIndexes]
				} else {
					score -= p.scores[tmpIndexes]
				}
			}
		}
	}
	if board.GetCellColorCount(!color) > board.GetCellColorCount(color) {
		score += score / 2
	}

	return score
}

func (p *MinimaxBot) minGameplay(board game.Board, color game.Color, depth int) int {
	if depth == 0 {
		return p.calculateHeuristics(board, color, depth)
	}
	turns := board.GetAllTurns(!color)
	if len(turns) == 0 {
		return p.maxGameplay(board, color, depth-1)
	}

	heurScore := MAX_SCORE

	for _, turn := range turns {
		newBoard := board.Copy()
		err := newBoard.UpdateBoard(turn)
		if err != nil {
			fmt.Println("ERROR IN minGameplay TURN", err)
		}
		score := p.maxGameplay(newBoard, color, depth-1)
		if score < heurScore {
			heurScore = score
		}

	}

	return heurScore
}
func (p *MinimaxBot) maxGameplay(board game.Board, color game.Color, depth int) int {
	if depth == 0 {
		return p.calculateHeuristics(board, color, depth)
	}

	turns := board.GetAllTurns(color)
	if len(turns) == 0 {
		return p.minGameplay(board, color, depth-1)
	}

	heurScore := -MAX_SCORE

	for _, turn := range turns {
		newBoard := board.Copy()
		err := newBoard.UpdateBoard(turn)
		if err != nil {
			fmt.Println("ERROR IN maxGameplay TURN", err)
		}
		score := p.maxGameplay(newBoard, color, depth-1)
		if score > heurScore {
			heurScore = score
		}

	}

	return heurScore
}
