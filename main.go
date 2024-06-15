package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/luispacheco-dev/snake-go/game"
)

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Snake-Go")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
