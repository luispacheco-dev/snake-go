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
		head:     head,
		tail:     head,
		width:    10,
		height:   10,
		deltaX:   0,
		deltaY:   0,
		nextPosX: head.x,
		nextPosY: head.y,
	}
}

func (s *Snake) Update() {
	speed := s.width / (float32(ebiten.TPS()) / 10)

	if s.deltaX == 0 && s.deltaY == 0 {
		s.deltaX = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && s.deltaY == 0 {
		s.deltaX = 0
		s.deltaY = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && s.deltaY == 0 {
		s.deltaX = 0
		s.deltaY = +speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && s.deltaX == 0 {
		s.deltaX = -speed
		s.deltaY = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && s.deltaX == 0 {
		s.deltaX = +speed
		s.deltaY = 0
	}

	s.nextPosX += s.deltaX
	s.nextPosY += s.deltaY

	s.move()
}

func (s *Snake) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, s.head.x, s.head.y, s.width, s.height, color.Black, false)
	snakePortion := s.head
	for snakePortion.nextPortion != nil {
		vector.DrawFilledRect(screen, snakePortion.x, snakePortion.y, s.width, s.height, color.Black, false)
		snakePortion = snakePortion.nextPortion
		if snakePortion == s.tail {
			vector.DrawFilledRect(screen, snakePortion.x, snakePortion.y, s.width, s.height, color.Black, false)
		}
	}
}

// ---

func (s *Snake) Grow() {
	s.changeHead()
}

// ---

func (s *Snake) move() {
	s.changeHead()
	s.removeTail()
}

func (s *Snake) changeHead() {
	snakePortion := &SnakePortion{
		x: s.nextPosX,
		y: s.nextPosY,
	}

	s.head.prevPortion = snakePortion
	snakePortion.nextPortion = s.head

	s.head = snakePortion
}

func (s *Snake) removeTail() {
	snakePortion := s.tail.prevPortion

	s.tail.prevPortion = nil
	snakePortion.nextPortion = nil

	s.tail = snakePortion
}
