package main

import (
	"Reversi/controller"
	"Reversi/game"
	"Reversi/view"
	"math/rand"
	"time"
)

func StartGame(gameController controller.Controller, gameView view.View) {
	events := make(chan game.Event)
	go gameController.Run(events)
	gameView.Run(events)
}

func run() {
	rand.Seed(time.Now().UTC().UnixNano())
	gameView := view.NewGameView()
	gameController := controller.NewGameRunner()
	StartGame(gameController, gameView)
}

func main() {
	run()
}
