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
	food  *Food
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

	g.food = NewFood(g)
	g.snake = NewSnake(g)
	return g
}

func (g *Game) Update() error {
	if g.checkCollisionSnakeBoard() {
		return fmt.Errorf("Snake hit board limits.")
	}
	if g.checkCollisionSnakeSnake() {
		return fmt.Errorf("Snake bit its own body/tail.")
	}
	if g.checkCollisionSnakeFood() {
		g.score++
		g.snake.Grow()
		g.food.randomXY(g.board)
	}
	g.snake.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("SCORE: %d", g.score))

	clr := color.RGBA{161, 195, 152, 0}
	vector.DrawFilledRect(screen, g.board.x, g.board.y, g.board.width, g.board.height, clr, false)

	g.food.Draw(screen)
	g.snake.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// ---

func (g *Game) checkCollisionSnakeBoard() bool {
	bx := g.board.x
	bMaxX := bx + g.board.width

	by := g.board.y
	bMaxY := by + g.board.height

	sx := g.snake.head.x
	sy := g.snake.head.y

	return sx <= bx || (sx+g.snake.width) >= bMaxX || sy <= by || (sy+g.snake.height) >= bMaxY
}

func (g *Game) checkCollisionSnakeFood() bool {
	sx := g.snake.head.x
	sMaxX := sx + g.snake.width

	sy := g.snake.head.y
	sMaxY := sy + g.snake.height

	fx := g.food.x
	fMaxX := fx + g.food.width

	fy := g.food.y
	fMaxY := fy + g.food.height

	return sx <= fMaxX && sMaxX >= fx && sy <= fMaxY && sMaxY >= fy
}

func (g *Game) checkCollisionSnakeSnake() bool {
	if g.snake.head == g.snake.tail {
		return false
	}

	sx := g.snake.head.x
	sy := g.snake.head.y

	snakePortion := g.snake.head.nextPortion
	for snakePortion.nextPortion != nil {
		spx := snakePortion.x
		spy := snakePortion.y
		if sx == spx && sy == spy {
			return true
		}
		snakePortion = snakePortion.nextPortion
		if snakePortion == g.snake.tail {
			return sx == snakePortion.x && sy == snakePortion.y
		}
	}

	return false
}
