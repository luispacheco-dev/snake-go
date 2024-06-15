package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	score int
	board Rect
	snake *Snake
}

func NewGame() *Game {
	board := Rect{
		x:      0,
		y:      0,
		width:  400,
		height: 400,
	}

	g := &Game{
		score: 0,
		board: board,
	}

	g.snake = NewSnake(g)
	return g
}

func (g *Game) Update() error {
	g.snake.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("SCORE: %d", g.score))

	clr := color.RGBA{161, 195, 152, 0}
	vector.DrawFilledRect(screen, g.board.x, g.board.y, g.board.width, g.board.height, clr, false)

	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
