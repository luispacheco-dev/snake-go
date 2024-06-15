package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	head     *SnakePortion
	tail     *SnakePortion
	width    float32
	height   float32
	deltaX   float32
	deltaY   float32
	nextPosX float32
	nextPosY float32
}

type SnakePortion struct {
	x           float32
	y           float32
	nextPortion *SnakePortion
	prevPortion *SnakePortion
}

func NewSnake(g *Game) *Snake {
	head := &SnakePortion{
		x:           g.board.width / 2,
		y:           g.board.height / 2,
		nextPortion: nil,
		prevPortion: nil,
	}

	return &Snake{
		head:   head,
		tail:   head,
		width:  10,
		height: 10,
		deltaX: 0,
		deltaY: 0,
	}
}

func (s *Snake) Update() {}

func (s *Snake) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.head.x, s.head.y, s.width, s.height, color.Black, false)
}

// ---

func (s *Snake) Grow() {}

func (s *Snake) Move() {}

// ---

func (s *Snake) changeHead() {}

func (s *Snake) removeTail() {}
